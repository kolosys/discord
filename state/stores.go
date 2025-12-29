package state

import (
	"context"
	"fmt"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/rest"
)

// GuildStore manages guild cache with relationship methods.
type GuildStore struct {
	*Store[string, *CachedGuild]
	rest     *rest.REST
	channels *ChannelStore
	roles    *RoleStore
	members  *MemberStore
}

func (s *GuildStore) Refresh(ctx context.Context, id string) (*CachedGuild, error) {
	if s.rest == nil {
		return nil, ErrRESTNotAvailable
	}
	if !s.Enabled() {
		return nil, ErrDisabled
	}
	guild, err := s.rest.GetGuild(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetch guild: %w", err)
	}
	cached := guildWithCountsToCached(guild)
	s.Set(ctx, id, cached)
	return cached, nil
}

func (s *GuildStore) Channels(ctx context.Context, guildID string) ([]*models.GuildChannel, error) {
	if !s.Enabled() || !s.channels.Enabled() {
		return nil, ErrDisabled
	}
	guild, found := s.Get(ctx, guildID)
	if !found {
		return nil, ErrNotFound
	}
	result := make([]*models.GuildChannel, 0, len(guild.ChannelIDs))
	for _, id := range guild.ChannelIDs {
		if ch, ok := s.channels.Get(ctx, id); ok {
			result = append(result, ch)
		}
	}
	return result, nil
}

func (s *GuildStore) Roles(ctx context.Context, guildID string) ([]*models.GuildRole, error) {
	if !s.Enabled() || !s.roles.Enabled() {
		return nil, ErrDisabled
	}
	guild, found := s.Get(ctx, guildID)
	if !found {
		return nil, ErrNotFound
	}
	result := make([]*models.GuildRole, 0, len(guild.RoleIDs))
	for _, id := range guild.RoleIDs {
		if role, ok := s.roles.Get(ctx, guildID, id); ok {
			result = append(result, role)
		}
	}
	return result, nil
}

func (s *GuildStore) Members(ctx context.Context, guildID string) ([]*CachedMember, error) {
	if !s.Enabled() || !s.members.Enabled() {
		return nil, ErrDisabled
	}
	guild, found := s.Get(ctx, guildID)
	if !found {
		return nil, ErrNotFound
	}
	result := make([]*CachedMember, 0, len(guild.MemberIDs))
	for _, userID := range guild.MemberIDs {
		if m, ok := s.members.Get(ctx, guildID, userID); ok {
			result = append(result, m)
		}
	}
	return result, nil
}

// ChannelStore manages channel cache.
type ChannelStore struct {
	*Store[string, *models.GuildChannel]
	rest *rest.REST
}

func (s *ChannelStore) Refresh(ctx context.Context, id string) (*models.GuildChannel, error) {
	if s.rest == nil {
		return nil, ErrRESTNotAvailable
	}
	if !s.Enabled() {
		return nil, ErrDisabled
	}
	chAny, err := s.rest.GetChannel(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetch channel: %w", err)
	}
	ch, ok := chAny.(*models.GuildChannel)
	if !ok {
		return nil, fmt.Errorf("unexpected channel type: %T", chAny)
	}
	s.Set(ctx, id, ch)
	return ch, nil
}

// UserStore manages user cache.
type UserStore struct {
	*Store[string, *models.User]
	rest *rest.REST
}

func (s *UserStore) Refresh(ctx context.Context, id string) (*models.User, error) {
	if s.rest == nil {
		return nil, ErrRESTNotAvailable
	}
	if !s.Enabled() {
		return nil, ErrDisabled
	}
	user, err := s.rest.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetch user: %w", err)
	}
	s.Set(ctx, id, user)
	return user, nil
}

// MemberStore manages member cache with composite keys.
type MemberStore struct {
	*Store[string, *CachedMember]
	rest *rest.REST
}

func (s *MemberStore) Get(ctx context.Context, guildID, userID string) (*CachedMember, bool) {
	return s.Store.Get(ctx, memberKey(guildID, userID))
}

func (s *MemberStore) Set(ctx context.Context, guildID, userID string, member *models.GuildMember) {
	if s == nil || !s.Enabled() || member == nil {
		return
	}
	cached := &CachedMember{
		GuildMember: *member,
		GuildID:     guildID,
		UserID:      userID,
	}
	s.Store.Set(ctx, memberKey(guildID, userID), cached)
}

func (s *MemberStore) Delete(ctx context.Context, guildID, userID string) bool {
	return s.Store.Delete(ctx, memberKey(guildID, userID))
}

func (s *MemberStore) Refresh(ctx context.Context, guildID, userID string) (*CachedMember, error) {
	if s.rest == nil {
		return nil, ErrRESTNotAvailable
	}
	if !s.Enabled() {
		return nil, ErrDisabled
	}
	member, err := s.rest.GetGuildMember(ctx, guildID, userID)
	if err != nil {
		return nil, fmt.Errorf("fetch member: %w", err)
	}
	cached := &CachedMember{
		GuildMember: *member,
		GuildID:     guildID,
		UserID:      userID,
	}
	s.Store.Set(ctx, memberKey(guildID, userID), cached)
	return cached, nil
}

// RoleStore manages role cache with composite keys.
type RoleStore struct {
	*Store[string, *models.GuildRole]
	rest *rest.REST
}

func (s *RoleStore) Get(ctx context.Context, guildID, roleID string) (*models.GuildRole, bool) {
	return s.Store.Get(ctx, roleKey(guildID, roleID))
}

func (s *RoleStore) Set(ctx context.Context, guildID string, role *models.GuildRole) {
	if s == nil || !s.Enabled() || role == nil {
		return
	}
	s.Store.Set(ctx, roleKey(guildID, role.ID), role)
}

func (s *RoleStore) Delete(ctx context.Context, guildID, roleID string) bool {
	return s.Store.Delete(ctx, roleKey(guildID, roleID))
}

// PresenceStore manages presence cache.
type PresenceStore struct {
	*Store[string, *Presence]
}

func (s *PresenceStore) Get(ctx context.Context, guildID, userID string) (*Presence, bool) {
	return s.Store.Get(ctx, presenceKey(guildID, userID))
}

func (s *PresenceStore) Set(ctx context.Context, presence *Presence) {
	if s == nil || !s.Enabled() || presence == nil {
		return
	}
	s.Store.Set(ctx, presenceKey(presence.GuildID, presence.UserID), presence)
}

func (s *PresenceStore) Delete(ctx context.Context, guildID, userID string) bool {
	return s.Store.Delete(ctx, presenceKey(guildID, userID))
}

// VoiceStateStore manages voice state cache.
type VoiceStateStore struct {
	*Store[string, *VoiceState]
}

func (s *VoiceStateStore) Get(ctx context.Context, guildID, userID string) (*VoiceState, bool) {
	return s.Store.Get(ctx, voiceStateKey(guildID, userID))
}

func (s *VoiceStateStore) Set(ctx context.Context, vs *VoiceState) {
	if s == nil || !s.Enabled() || vs == nil {
		return
	}
	s.Store.Set(ctx, voiceStateKey(vs.GuildID, vs.UserID), vs)
}

func (s *VoiceStateStore) Delete(ctx context.Context, guildID, userID string) bool {
	return s.Store.Delete(ctx, voiceStateKey(guildID, userID))
}

// guildWithCountsToCached converts a GuildWithCounts to CachedGuild.
func guildWithCountsToCached(g *models.GuildWithCounts) *CachedGuild {
	return &CachedGuild{
		Guild: models.Guild{
			ID:                          g.ID,
			AfkChannelID:                g.AfkChannelID,
			AfkTimeout:                  g.AfkTimeout,
			ApplicationID:               g.ApplicationID,
			Banner:                      g.Banner,
			DefaultMessageNotifications: g.DefaultMessageNotifications,
			Description:                 g.Description,
			DiscoverySplash:             g.DiscoverySplash,
			Emojis:                      g.Emojis,
			ExplicitContentFilter:       g.ExplicitContentFilter,
			Features:                    g.Features,
			HomeHeader:                  g.HomeHeader,
			Icon:                        g.Icon,
			MaxMembers:                  g.MaxMembers,
			MaxPresences:                g.MaxPresences,
			MaxStageVideoChannelUsers:   g.MaxStageVideoChannelUsers,
			MaxVideoChannelUsers:        g.MaxVideoChannelUsers,
			MfaLevel:                    g.MfaLevel,
			Name:                        g.Name,
			Nsfw:                        g.Nsfw,
			NsfwLevel:                   g.NsfwLevel,
			OwnerID:                     g.OwnerID,
			PreferredLocale:             g.PreferredLocale,
			PremiumProgressBarEnabled:   g.PremiumProgressBarEnabled,
			PremiumSubscriptionCount:    g.PremiumSubscriptionCount,
			PremiumTier:                 g.PremiumTier,
			PublicUpdatesChannelID:      g.PublicUpdatesChannelID,
			Region:                      g.Region,
			Roles:                       g.Roles,
			RulesChannelID:              g.RulesChannelID,
			SafetyAlertsChannelID:       g.SafetyAlertsChannelID,
			Splash:                      g.Splash,
			Stickers:                    g.Stickers,
			SystemChannelFlags:          g.SystemChannelFlags,
			SystemChannelID:             g.SystemChannelID,
			VanityURLCode:               g.VanityURLCode,
			VerificationLevel:           g.VerificationLevel,
			WidgetChannelID:             g.WidgetChannelID,
			WidgetEnabled:               g.WidgetEnabled,
		},
		MemberCount: int(derefInt32(g.ApproximateMemberCount)),
	}
}

func derefInt32(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}
