package types_test

import (
	"testing"

	"github.com/kolosys/discord/types"
)

func TestChannelType_IsThread(t *testing.T) {
	tests := []struct {
		name         string
		channelType  types.ChannelType
		expectedTrue bool
	}{
		{"Text channel is not thread", types.ChannelTypeGuildText, false},
		{"DM is not thread", types.ChannelTypeDM, false},
		{"Voice is not thread", types.ChannelTypeGuildVoice, false},
		{"Group DM is not thread", types.ChannelTypeGroupDM, false},
		{"Category is not thread", types.ChannelTypeGuildCategory, false},
		{"Announcement is not thread", types.ChannelTypeGuildAnnouncement, false},
		{"Announcement thread is thread", types.ChannelTypeAnnouncementThread, true},
		{"Public thread is thread", types.ChannelTypePublicThread, true},
		{"Private thread is thread", types.ChannelTypePrivateThread, true},
		{"Stage voice is not thread", types.ChannelTypeGuildStageVoice, false},
		{"Guild directory is not thread", types.ChannelTypeGuildDirectory, false},
		{"Forum is not thread", types.ChannelTypeGuildForum, false},
		{"Media is not thread", types.ChannelTypeGuildMedia, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.channelType.IsThread()
			if got != tt.expectedTrue {
				t.Errorf("IsThread() = %v, want %v", got, tt.expectedTrue)
			}
		})
	}
}

func TestChannelType_IsVoice(t *testing.T) {
	tests := []struct {
		name         string
		channelType  types.ChannelType
		expectedTrue bool
	}{
		{"Text channel is not voice", types.ChannelTypeGuildText, false},
		{"DM is not voice", types.ChannelTypeDM, false},
		{"Voice is voice", types.ChannelTypeGuildVoice, true},
		{"Group DM is not voice", types.ChannelTypeGroupDM, false},
		{"Category is not voice", types.ChannelTypeGuildCategory, false},
		{"Announcement is not voice", types.ChannelTypeGuildAnnouncement, false},
		{"Announcement thread is not voice", types.ChannelTypeAnnouncementThread, false},
		{"Public thread is not voice", types.ChannelTypePublicThread, false},
		{"Private thread is not voice", types.ChannelTypePrivateThread, false},
		{"Stage voice is voice", types.ChannelTypeGuildStageVoice, true},
		{"Guild directory is not voice", types.ChannelTypeGuildDirectory, false},
		{"Forum is not voice", types.ChannelTypeGuildForum, false},
		{"Media is not voice", types.ChannelTypeGuildMedia, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.channelType.IsVoice()
			if got != tt.expectedTrue {
				t.Errorf("IsVoice() = %v, want %v", got, tt.expectedTrue)
			}
		})
	}
}

func TestChannelType_IsGuild(t *testing.T) {
	tests := []struct {
		name         string
		channelType  types.ChannelType
		expectedTrue bool
	}{
		{"Text channel is guild", types.ChannelTypeGuildText, true},
		{"DM is not guild", types.ChannelTypeDM, false},
		{"Voice is guild", types.ChannelTypeGuildVoice, true},
		{"Group DM is not guild", types.ChannelTypeGroupDM, false},
		{"Category is guild", types.ChannelTypeGuildCategory, true},
		{"Announcement is guild", types.ChannelTypeGuildAnnouncement, true},
		{"Announcement thread is guild", types.ChannelTypeAnnouncementThread, true},
		{"Public thread is guild", types.ChannelTypePublicThread, true},
		{"Private thread is guild", types.ChannelTypePrivateThread, true},
		{"Stage voice is guild", types.ChannelTypeGuildStageVoice, true},
		{"Guild directory is guild", types.ChannelTypeGuildDirectory, true},
		{"Forum is guild", types.ChannelTypeGuildForum, true},
		{"Media is guild", types.ChannelTypeGuildMedia, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.channelType.IsGuild()
			if got != tt.expectedTrue {
				t.Errorf("IsGuild() = %v, want %v", got, tt.expectedTrue)
			}
		})
	}
}

func TestChannelType_Combinations(t *testing.T) {
	tests := []struct {
		name           string
		channelType    types.ChannelType
		shouldBeThread bool
		shouldBeVoice  bool
		shouldBeGuild  bool
	}{
		{"Announcement thread", types.ChannelTypeAnnouncementThread, true, false, true},
		{"Public thread", types.ChannelTypePublicThread, true, false, true},
		{"Private thread", types.ChannelTypePrivateThread, true, false, true},
		{"Voice channel", types.ChannelTypeGuildVoice, false, true, true},
		{"Stage voice", types.ChannelTypeGuildStageVoice, false, true, true},
		{"Text channel", types.ChannelTypeGuildText, false, false, true},
		{"DM", types.ChannelTypeDM, false, false, false},
		{"Group DM", types.ChannelTypeGroupDM, false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isThread := tt.channelType.IsThread(); isThread != tt.shouldBeThread {
				t.Errorf("IsThread() = %v, want %v", isThread, tt.shouldBeThread)
			}
			if isVoice := tt.channelType.IsVoice(); isVoice != tt.shouldBeVoice {
				t.Errorf("IsVoice() = %v, want %v", isVoice, tt.shouldBeVoice)
			}
			if isGuild := tt.channelType.IsGuild(); isGuild != tt.shouldBeGuild {
				t.Errorf("IsGuild() = %v, want %v", isGuild, tt.shouldBeGuild)
			}
		})
	}
}
