package models

type Tag struct {
	Name       string   `bson:"_id" json:"name"`                  // Tag Name that appears
	BlogIDList []string `bson:"blog_id_list" json:"blog_id_list"` // List of BlogIDs associated with the Tag
}
