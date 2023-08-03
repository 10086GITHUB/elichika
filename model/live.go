package model

// SaveDeckReq ...
type SaveDeckReq struct {
	DeckID       int   `json:"deck_id"`
	CardWithSuit []int `json:"card_with_suit"`
	SquadDict    []any `json:"squad_dict"`
}

// UserLiveDeck ...
type UserLiveDeck struct {
	UserID         int `xorm:"pk 'user_id'" json:"-"`
	UserLiveDeckID int `xorm:"pk 'user_live_deck_id'" json:"user_live_deck_id"`
	Name           struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // deck name
	CardMasterID1 int `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterID2 int `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterID3 int `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	CardMasterID4 int `xorm:"'card_master_id_4'" json:"card_master_id_4"`
	CardMasterID5 int `xorm:"'card_master_id_5'" json:"card_master_id_5"`
	CardMasterID6 int `xorm:"'card_master_id_6'" json:"card_master_id_6"`
	CardMasterID7 int `xorm:"'card_master_id_7'" json:"card_master_id_7"`
	CardMasterID8 int `xorm:"'card_master_id_8'" json:"card_master_id_8"`
	CardMasterID9 int `xorm:"'card_master_id_9'" json:"card_master_id_9"`
	SuitMasterID1 int `xorm:"'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterID2 int `xorm:"'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterID3 int `xorm:"'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterID4 int `xorm:"'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterID5 int `xorm:"'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterID6 int `xorm:"'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterID7 int `xorm:"'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterID8 int `xorm:"'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterID9 int `xorm:"'suit_master_id_9'" json:"suit_master_id_9"`
}

// UserLiveParty ...
type UserLiveParty struct {
	UserID         int `xorm:"pk 'user_id'" json:"-"`
	PartyID        int `xorm:"pk 'party_id'" json:"party_id"`
	UserLiveDeckID int `xorm:"'user_live_deck_id'" json:"user_live_deck_id"`
	Name           struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // deck name
	IconMasterID     int    `xorm:"'icon_master_id'" json:"icon_master_id"`
	CardMasterID1    int    `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterID2    int    `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterID3    int    `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	UserAccessoryID1 *int64 `xorm:"'user_accessory_id_1'" json:"user_accessory_id_1"` // null for empty
	UserAccessoryID2 *int64 `xorm:"'user_accessory_id_2'" json:"user_accessory_id_2"`
	UserAccessoryID3 *int64 `xorm:"'user_accessory_id_3'" json:"user_accessory_id_3"`
}

// PartyName ...
type PartyName struct {
	DotUnderText string `json:"dot_under_text"`
}

// DeckSquadDict ...
type DeckSquadDict struct {
	CardMasterIds    []int    `json:"card_master_ids"`
	UserAccessoryIds []*int64 `json:"user_accessory_ids"`
}

// LiveDaily ...
type LiveDaily struct {
	LiveDailyMasterID      int `json:"live_daily_master_id" xorm:"id"`
	LiveMasterID           int `json:"live_master_id" xorm:"live_id"`
	EndAt                  int `json:"end_at"`
	RemainingPlayCount     int `json:"remaining_play_count"`
	RemainingRecoveryCount int `json:"remaining_recovery_count"`
}

// LiveStartReq ...
type LiveStartReq struct {
	LiveDifficultyID    int  `json:"live_difficulty_id"`
	DeckID              int  `json:"deck_id"`
	PartnerUserID       int  `json:"partner_user_id"`
	PartnerCardMasterID int  `json:"partner_card_master_id"`
	LpMagnification     int  `json:"lp_magnification"`
	IsAutoPlay          bool `json:"is_auto_play"`
	IsReferenceBook     bool `json:"is_reference_book"`
}

// LivePartnerInfo ...
type LivePartnerInfo struct {
	UserID                              int                 `json:"user_id"`
	Name                                PartnerName         `json:"name"`
	Rank                                int                 `json:"rank"`
	LastPlayedAt                        int64               `json:"last_played_at"`
	RecommendCardMasterID               int                 `json:"recommend_card_master_id"`
	RecommendCardLevel                  int                 `json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool                `json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool                `json:"is_recommend_card_all_training_activated"`
	EmblemID                            int                 `json:"emblem_id"`
	IsNew                               bool                `json:"is_new"`
	IntroductionMessage                 IntroductionMessage `json:"introduction_message"`
	FriendApprovedAt                    any                 `json:"friend_approved_at"`
	RequestStatus                       int                 `json:"request_status"`
	IsRequestPending                    bool                `json:"is_request_pending"`
}

// PartnerName ...
type PartnerName struct {
	DotUnderText string `json:"dot_under_text"`
}

// IntroductionMessage ...
type IntroductionMessage struct {
	DotUnderText string `json:"dot_under_text"`
}

// LiveResultAchievementStatus ...
type LiveResultAchievementStatus struct {
	ClearCount       int64 `json:"clear_count"`
	GotVoltage       int64 `json:"got_voltage"`
	RemainingStamina int64 `json:"remaining_stamina"`
}

// MvpInfo ...
type MvpInfo struct {
	CardMasterID        int64 `json:"card_master_id"`
	GetVoltage          int64 `json:"get_voltage"`
	SkillTriggeredCount int64 `json:"skill_triggered_count"`
	AppealCount         int64 `json:"appeal_count"`
}

// LiveSaveDeckReq ...
type LiveSaveDeckReq struct {
	LiveMasterID        int   `json:"live_master_id"`
	LiveMvDeckType      int   `json:"live_mv_deck_type"`
	MemberMasterIDByPos []int `json:"member_master_id_by_pos"`
	SuitMasterIDByPos   []int `json:"suit_master_id_by_pos"`
	ViewStatusByPos     []int `json:"view_status_by_pos"`
}

// UserLiveMvDeckInfo ...
type UserLiveMvDeckInfo struct {
	LiveMasterID     any `json:"live_master_id"`
	MemberMasterID1  any `json:"member_master_id_1"`
	MemberMasterID2  any `json:"member_master_id_2"`
	MemberMasterID3  any `json:"member_master_id_3"`
	MemberMasterID4  any `json:"member_master_id_4"`
	MemberMasterID5  any `json:"member_master_id_5"`
	MemberMasterID6  any `json:"member_master_id_6"`
	MemberMasterID7  any `json:"member_master_id_7"`
	MemberMasterID8  any `json:"member_master_id_8"`
	MemberMasterID9  any `json:"member_master_id_9"`
	MemberMasterID10 any `json:"member_master_id_10"`
	MemberMasterID11 any `json:"member_master_id_11"`
	MemberMasterID12 any `json:"member_master_id_12"`
	SuitMasterID1    any `json:"suit_master_id_1"`
	SuitMasterID2    any `json:"suit_master_id_2"`
	SuitMasterID3    any `json:"suit_master_id_3"`
	SuitMasterID4    any `json:"suit_master_id_4"`
	SuitMasterID5    any `json:"suit_master_id_5"`
	SuitMasterID6    any `json:"suit_master_id_6"`
	SuitMasterID7    any `json:"suit_master_id_7"`
	SuitMasterID8    any `json:"suit_master_id_8"`
	SuitMasterID9    any `json:"suit_master_id_9"`
	SuitMasterID10   any `json:"suit_master_id_10"`
	SuitMasterID11   any `json:"suit_master_id_11"`
	SuitMasterID12   any `json:"suit_master_id_12"`
}

// LiveStageInfo ...
type LiveStageInfo struct {
	LiveDifficultyID int                `json:"live_difficulty_id"`
	LiveNotes        []LiveNotes        `json:"live_notes"`
	LiveWaveSettings []LiveWaveSettings `json:"live_wave_settings"`
	NoteGimmicks     []NoteGimmicks     `json:"note_gimmicks"`
	StageGimmickDict []any              `json:"stage_gimmick_dict"`
}

// LiveNotes ...
type LiveNotes struct {
	ID                  int `json:"id"`
	CallTime            int `json:"call_time"`
	NoteType            int `json:"note_type"`
	NotePosition        int `json:"note_position"`
	GimmickID           int `json:"gimmick_id"`
	NoteAction          int `json:"note_action"`
	WaveID              int `json:"wave_id"`
	NoteRandomDropColor int `json:"note_random_drop_color"`
	AutoJudgeType       int `json:"auto_judge_type"`
}

// LiveWaveSettings ...
type LiveWaveSettings struct {
	ID            int `json:"id"`
	WaveDamage    int `json:"wave_damage"`
	MissionType   int `json:"mission_type"`
	Arg1          int `json:"arg_1"`
	Arg2          int `json:"arg_2"`
	RewardVoltage int `json:"reward_voltage"`
}

// NoteGimmicks ...
type NoteGimmicks struct {
	UniqID          int `json:"uniq_id"`
	ID              int `json:"id"`
	NoteGimmickType int `json:"note_gimmick_type"`
	Arg1            int `json:"arg_1"`
	Arg2            int `json:"arg_2"`
	EffectMID       int `json:"effect_m_id"`
	IconType        int `json:"icon_type"`
}

// PartnerCardReq ...
type PartnerCardReq struct {
	LivePartnerCategoryID int `json:"live_partner_category_id"`
	CardMasterID          int `json:"card_master_id"`
}

// PartnerCard ...
type PartnerCard struct {
	CardMasterID              int                `json:"card_master_id"`
	Level                     int                `json:"level"`
	Grade                     int                `json:"grade"`
	LoveLevel                 int                `json:"love_level"`
	IsAwakening               bool               `json:"is_awakening"`
	IsAwakeningImage          bool               `json:"is_awakening_image"`
	IsAllTrainingActivated    bool               `json:"is_all_training_activated"`
	ActiveSkillLevel          int                `json:"active_skill_level"`
	PassiveSkillLevels        []int              `json:"passive_skill_levels"`
	AdditionalPassiveSkillIds []int              `json:"additional_passive_skill_ids"`
	MaxFreePassiveSkill       int                `json:"max_free_passive_skill"`
	TrainingStamina           int                `json:"training_stamina"`
	TrainingAppeal            int                `json:"training_appeal"`
	TrainingTechnique         int                `json:"training_technique"`
	MemberLovePanels          []MemberLovePanels `json:"member_love_panels"`
}

// MemberLovePanels ...
type MemberLovePanels struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelCellIds []int `json:"member_love_panel_cell_ids"`
}
