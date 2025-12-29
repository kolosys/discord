package state

import (
	"context"

	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/models"
)

// RegisterHandlers sets up event handlers for automatic state population.
func (s *State) RegisterHandlers(dispatcher *events.Dispatcher) error {
	if err := events.On(dispatcher, events.Ready, s.handleReady); err != nil {
		return err
	}

	if s.Guilds.Enabled() {
		if err := events.On(dispatcher, events.GuildCreate, s.handleGuildCreate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildUpdate, s.handleGuildUpdate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildDelete, s.handleGuildDelete); err != nil {
			return err
		}
	}

	if s.Channels.Enabled() {
		if err := events.On(dispatcher, events.ChannelCreate, s.handleChannelCreate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.ChannelUpdate, s.handleChannelUpdate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.ChannelDelete, s.handleChannelDelete); err != nil {
			return err
		}
	}

	if s.Members.Enabled() {
		if err := events.On(dispatcher, events.GuildMemberAdd, s.handleMemberAdd); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildMemberUpdate, s.handleMemberUpdate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildMemberRemove, s.handleMemberRemove); err != nil {
			return err
		}
	}

	if s.Roles.Enabled() {
		if err := events.On(dispatcher, events.GuildRoleCreate, s.handleRoleCreate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildRoleUpdate, s.handleRoleUpdate); err != nil {
			return err
		}
		if err := events.On(dispatcher, events.GuildRoleDelete, s.handleRoleDelete); err != nil {
			return err
		}
	}

	if s.Presences.Enabled() {
		if err := events.On(dispatcher, events.PresenceUpdate, s.handlePresenceUpdate); err != nil {
			return err
		}
	}

	if s.VoiceStates.Enabled() {
		if err := events.On(dispatcher, events.VoiceStateUpdate, s.handleVoiceStateUpdate); err != nil {
			return err
		}
	}

	return nil
}

func (s *State) handleReady(ctx context.Context, e *events.ReadyEvent) {
	s.setCurrentUser(e.User)
	s.setReady(true)
}

func (s *State) handleGuildCreate(ctx context.Context, e *events.GuildCreateEvent) {
	if !s.Guilds.Enabled() {
		return
	}

	cached := &CachedGuild{
		Guild:       e.Guild,
		JoinedAt:    e.JoinedAt,
		Large:       e.Large,
		Unavailable: e.Unavailable,
		MemberCount: e.MemberCount,
	}

	if s.Channels.Enabled() {
		for i := range e.Channels {
			ch := e.Channels[i]
			s.Channels.Set(ctx, ch.ID, &ch)
			cached.ChannelIDs = append(cached.ChannelIDs, ch.ID)
		}
	}

	if s.Members.Enabled() {
		for i := range e.Members {
			m := e.Members[i]
			userID := extractUserID(&m)
			if userID != "" {
				s.Members.Set(ctx, e.Guild.ID, userID, &m)
				cached.MemberIDs = append(cached.MemberIDs, userID)
			}
		}
	}

	if s.Roles.Enabled() {
		for _, roleAny := range e.Guild.Roles {
			if role, ok := roleAny.(models.GuildRole); ok {
				s.Roles.Set(ctx, e.Guild.ID, &role)
				cached.RoleIDs = append(cached.RoleIDs, role.ID)
			}
		}
	}

	s.Guilds.Set(ctx, e.Guild.ID, cached)
}

func (s *State) handleGuildUpdate(ctx context.Context, e *events.GuildUpdateEvent) {
	if !s.Guilds.Enabled() {
		return
	}

	existing, found := s.Guilds.Get(ctx, e.Guild.ID)
	if found {
		existing.Guild = e.Guild
		s.Guilds.Set(ctx, e.Guild.ID, existing)
	} else {
		s.Guilds.Set(ctx, e.Guild.ID, &CachedGuild{Guild: e.Guild})
	}
}

func (s *State) handleGuildDelete(ctx context.Context, e *events.GuildDeleteEvent) {
	if e.Unavailable {
		if s.Guilds.Enabled() {
			if existing, found := s.Guilds.Get(ctx, e.ID); found {
				existing.Unavailable = true
				s.Guilds.Set(ctx, e.ID, existing)
			}
		}
		return
	}

	if s.Guilds.Enabled() {
		s.Guilds.Delete(ctx, e.ID)
	}
}

func (s *State) handleChannelCreate(ctx context.Context, e *events.ChannelCreateEvent) {
	if !s.Channels.Enabled() {
		return
	}
	s.Channels.Set(ctx, e.GuildChannel.ID, &e.GuildChannel)
}

func (s *State) handleChannelUpdate(ctx context.Context, e *events.ChannelUpdateEvent) {
	if !s.Channels.Enabled() {
		return
	}
	s.Channels.Set(ctx, e.GuildChannel.ID, &e.GuildChannel)
}

func (s *State) handleChannelDelete(ctx context.Context, e *events.ChannelDeleteEvent) {
	if !s.Channels.Enabled() {
		return
	}
	s.Channels.Delete(ctx, e.GuildChannel.ID)
}

func (s *State) handleMemberAdd(ctx context.Context, e *events.GuildMemberAddEvent) {
	if !s.Members.Enabled() {
		return
	}
	userID := extractUserID(&e.GuildMember)
	if userID != "" {
		s.Members.Set(ctx, e.GuildID, userID, &e.GuildMember)
	}

	if s.Guilds.Enabled() {
		if guild, found := s.Guilds.Get(ctx, e.GuildID); found {
			guild.MemberCount++
			if userID != "" {
				guild.MemberIDs = append(guild.MemberIDs, userID)
			}
			s.Guilds.Set(ctx, e.GuildID, guild)
		}
	}
}

func (s *State) handleMemberUpdate(ctx context.Context, e *events.GuildMemberUpdateEvent) {
	if !s.Members.Enabled() {
		return
	}
	userID := e.User.ID
	if userID != "" {
		member := &models.GuildMember{
			Deaf:         e.Deaf,
			JoinedAt:     e.JoinedAt,
			Mute:         e.Mute,
			Nick:         e.Nick,
			Pending:      e.Pending,
			PremiumSince: e.PremiumSince,
			Roles:        e.Roles,
		}
		s.Members.Set(ctx, e.GuildID, userID, member)
	}
}

func (s *State) handleMemberRemove(ctx context.Context, e *events.GuildMemberRemoveEvent) {
	if !s.Members.Enabled() {
		return
	}
	s.Members.Delete(ctx, e.GuildID, e.User.ID)

	if s.Guilds.Enabled() {
		if guild, found := s.Guilds.Get(ctx, e.GuildID); found {
			guild.MemberCount--
			if guild.MemberCount < 0 {
				guild.MemberCount = 0
			}
			for i, id := range guild.MemberIDs {
				if id == e.User.ID {
					guild.MemberIDs = append(guild.MemberIDs[:i], guild.MemberIDs[i+1:]...)
					break
				}
			}
			s.Guilds.Set(ctx, e.GuildID, guild)
		}
	}
}

func (s *State) handleRoleCreate(ctx context.Context, e *events.GuildRoleCreateEvent) {
	if !s.Roles.Enabled() {
		return
	}
	s.Roles.Set(ctx, e.GuildID, &e.Role)
}

func (s *State) handleRoleUpdate(ctx context.Context, e *events.GuildRoleUpdateEvent) {
	if !s.Roles.Enabled() {
		return
	}
	s.Roles.Set(ctx, e.GuildID, &e.Role)
}

func (s *State) handleRoleDelete(ctx context.Context, e *events.GuildRoleDeleteEvent) {
	if !s.Roles.Enabled() {
		return
	}
	s.Roles.Delete(ctx, e.GuildID, e.RoleID)
}

func (s *State) handlePresenceUpdate(ctx context.Context, e *events.PresenceUpdateEvent) {
	if !s.Presences.Enabled() {
		return
	}
	presence := &Presence{
		UserID:       e.User.ID,
		GuildID:      e.GuildID,
		Status:       e.Status,
		Activities:   e.Activities,
		ClientStatus: e.ClientStatus,
	}
	s.Presences.Set(ctx, presence)
}

func (s *State) handleVoiceStateUpdate(ctx context.Context, e *events.VoiceStateUpdateEvent) {
	if !s.VoiceStates.Enabled() {
		return
	}

	if e.ChannelID == nil || *e.ChannelID == "" {
		s.VoiceStates.Delete(ctx, e.GuildID, e.UserID)
		return
	}

	voiceState := &VoiceState{
		UserID:     e.UserID,
		GuildID:    e.GuildID,
		ChannelID:  e.ChannelID,
		SessionID:  e.SessionID,
		Deaf:       e.Deaf,
		Mute:       e.Mute,
		SelfDeaf:   e.SelfDeaf,
		SelfMute:   e.SelfMute,
		SelfStream: e.SelfStream,
		SelfVideo:  e.SelfVideo,
		Suppress:   e.Suppress,
	}
	s.VoiceStates.Set(ctx, voiceState)
}

func extractUserID(m *models.GuildMember) string {
	if m.User == nil {
		return ""
	}
	if u, ok := m.User.(*models.User); ok {
		return u.ID
	}
	if um, ok := m.User.(map[string]any); ok {
		if id, ok := um["id"].(string); ok {
			return id
		}
	}
	return ""
}
