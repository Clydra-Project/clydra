package pipelineimpl

import "time"

type timeStamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Pipeline struct {
	ID   int64  `gorm:"primary_key;auto_increment;not null"`
	UUID string `gorm:"type:varchar(36);not null;unique_index"`
	Name string `gorm:"type:varchar(255);not null"`
	timeStamps
}

type PipelineNodeType struct {
	ID   int64  `gorm:"primary_key;auto_increment;not null"`
	UUID string `gorm:"type:varchar(36);not null;unique_index"`
	Name string `gorm:"type:varchar(255);not null"`
	timeStamps
}

type PipelineNode struct {
	ID         int64            `gorm:"primary_key;auto_increment;not null"`
	UUID       string           `gorm:"type:varchar(36);not null;unique_index"`
	Name       string           `gorm:"type:varchar(255);not null"`
	Pipeline   Pipeline         `gorm:"foreignkey:PipelineID;association_foreignkey:ID"`
	PipelineID int64            `gorm:"not null"`
	NodeType   PipelineNodeType `gorm:"foreignkey:NodeTypeID;association_foreignkey:ID"`
	NodeTypeID int64            `gorm:"not null"`
	Config     string           `gorm:"type:text;not null"`
	timeStamps
}

type PipelineRunner struct {
	ID           int64         `gorm:"primary_key;auto_increment;not null"`
	UUID         string        `gorm:"type:varchar(36);not null;unique_index"`
	Name         string        `gorm:"type:varchar(255);not null"`
	PipelineID   int64         `gorm:"not null"`
	Pipeline     Pipeline      `gorm:"foreignkey:PipelineID;association_foreignkey:ID"`
	PipelineLogs []PipelineLog `gorm:"foreignkey:RunnerID;association_foreignkey:ID"`
	timeStamps
}

type PipelineLog struct {
	ID       int64          `gorm:"primary_key;auto_increment;not null"`
	UUID     string         `gorm:"type:varchar(36);not null;unique_index"`
	Runner   PipelineRunner `gorm:"foreignkey:RunnerID;association_foreignkey:ID"`
	RunnerID int64          `gorm:"not null"`
	Node     PipelineNode   `gorm:"foreignkey:NodeID;association_foreignkey:ID"`
	NodeID   int64          `gorm:"not null"`
	timeStamps
}

type PipelineEdgeNode struct {
	ID           int64        `gorm:"primary_key;auto_increment;not null"`
	UUID         string       `gorm:"type:varchar(36);not null;unique_index"`
	Pipeline     Pipeline     `gorm:"foreignkey:PipelineID;association_foreignkey:ID"`
	PipelineID   int64        `gorm:"not null"`
	NodeInput    PipelineNode `gorm:"foreignkey:NodeInputID;association_foreignkey:ID"`
	NodeInputID  int64        `gorm:"not null"`
	NodeOutput   PipelineNode `gorm:"foreignkey:NodeOutputID;association_foreignkey:ID"`
	NodeOutputID int64        `gorm:"not null"`
	timeStamps
}
