package models

// Collection Courses for mongoDB

// Courses collection for mongoDB
type Courses struct {
	Course []Course `json:"courses" bson:"courses"`
}

// Missions collection for mongoDB
type Missions struct {
	Mission []Mission `json:"missions" bson:"missions"`
}

// Lobbies collection for mongoDB
type Lobbies struct {
	Lobby []Lobby `json:"lobbies" bson:"lobbies"`
}

// Systems collection for mongoDB
type Systems struct {
	System []System `json:"systems" bson:"systems"`
}

// DimondConversations collection for mongoDB
type DimondConversations struct {
	DimondConversation []DimondConversation `json:"dimond_conversations" bson:"dimond_conversations"`
}

// Course is a representation of courses as a list of missions.
type Course struct {
	Id          string   `json:"id" bson:"_id"`
	AuthorId    string   `json:"author_id" bson:"author_id"`
	Name        string   `json:"name" bson:"name"`
	CreatedAt   string   `json:"created_at" bson:"created_at"`
	MissionsIds []string `json:"missions_ids" bson:"missions_ids"`
}

// Mission is a representation of mission.
type Mission struct {
	Id            string     `json:"id" bson:"_id"`
	AuthorId      string     `json:"author_id" bson:"author_id"`
	CompanyId     string     `json:"company_id" bson:"company_id"`
	Title         string     `json:"title" bson:"title"`
	MethodologyId string     `json:"methodology_id" bson:"methodology_id"`
	ImageUrl      string     `json:"image_url" bson:"image_url"`
	Intro         []Intro    `json:"intro" bson:"intro"`
	Materials     []Material `json:"materials" bson:"materials"`
	Levels        []Level    `json:"levels" bson:"levels"`
}

// Intro is a representation of mission intro.
type Intro struct {
	Id        string          `json:"id" bson:"_id"`
	Title     string          `json:"title" bson:"title"`
	Content   string          `json:"content" bson:"content"`
	Type      string          `json:"type" bson:"type"`
	Links     []string        `json:"links" bson:"links"`
	Materials []IntroMaterial `json:"materials" bson:"materials"`
}

// Material is a representation of mission materials.
type Material struct {
	Id          string `json:"id" bson:"_id"`
	Filename    string `json:"filename" bson:"filename"`
	VectorCount int    `json:"vector_count" bson:"vector_count"`
}

// IntroMaterial is a representation of mission intro materials.
type IntroMaterial struct {
	Filename string `json:"filename" bson:"filename"`
	FileUrl  string `json:"file_url" bson:"file_url"`
}

// Level is a representation of mission levels.
type Level struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	CaseContent string `json:"case_content" bson:"case_content"`
	Type        string `json:"type" bson:"type"`
}

// Lobby is a representation of lobbies as a list of missions.
type Lobby struct {
	Id       string                     `json:"id" bson:"_id"`
	AuthorId string                     `json:"author_id" bson:"author_id"`
	Name     string                     `json:"name" bson:"name"`
	CourseId string                     `json:"course_id" bson:"course_id"`
	Missions map[string]MissionProgress `json:"missions" bson:"missions"`
}

// MissionProgress is a representation of mission progress.
type MissionProgress struct {
	Summary string      `json:"summary" bson:"summary"`
	Answers []Answer    `json:"answers" bson:"answers"`
	Users   []LobbyUser `json:"users" bson:"users"`
}

// Answer is a representation of answer.
type Answer struct {
	UserId        string           `json:"user_id" bson:"user_id"`
	LevelId       string           `json:"level_id" bson:"level_id"`
	NextLevelInfo string           `json:"next_level_info" bson:"next_level_info"`
	NumDiamonds   int              `json:"num_diamonds" bson:"num_diamonds"`
	Completed     bool             `json:"completed" bson:"completed"`
	Checked       bool             `json:"checked" bson:"checked"`
	History       []HistoryElement `json:"history" bson:"history"`
}

// HistoryElement is a representation of history element.
type HistoryElement struct {
	Role    string `json:"role" bson:"role"`
	Content string `json:"content" bson:"content"`
}

// LobbyUser is a representation of lobby user.
type LobbyUser struct {
	UserId             string `json:"user_id" bson:"userId"`
	Summary            string `json:"summary" bson:"summary"`
	TotalDiamonds      int    `json:"total_diamonds" bson:"total_diamonds"`
	ImageUrl           string `json:"image_url" bson:"image_url"`
	LastCompletedLevel int    `json:"last_completed_level" bson:"last_completed_level"`
}

// System is a representation of system.
type System struct {
	Id                         string                    `json:"id" bson:"_id"`
	LevelType                  string                    `json:"level_type" bson:"level_type" `
	DataClass                  string                    `json:"data_class" bson:"data_class"`
	AgentClass                 string                    `json:"agent_class" bson:"agent_class"`
	RequiresDiamondsToComplete bool                      `json:"requires_diamonds_to_complete" bson:"requires_diamonds_to_complete"`
	DataClassParameters        SystemDataClassParameters `json:"data_class_parameters" bson:"data_class_parameters"`
}

// SystemDataClassParameters is a representation of data class parameters.
type SystemDataClassParameters struct {
	FieldClass      string                `json:"field_class" bson:"field_class"`
	FieldParameters SystemFieldParameters `json:"field_parameters" bson:"field_parameters"`
}

// SystemFieldParameters is a representation of field parameters.
type SystemFieldParameters struct {
	Title     string `json:"title" bson:"title"`
	Delimiter string `json:"delimiter" bson:"delimiter"`
}

// DimondConversation is a representation of dimond conversation.
type DimondConversation struct {
	Id        string           `json:"id" bson:"_id"`
	UserId    string           `json:"user_id" bson:"user_id"`
	DiamondId string           `json:"diamond_id" bson:"diamond_id"`
	History   []HistoryElement `json:"history" bson:"history"`
}
