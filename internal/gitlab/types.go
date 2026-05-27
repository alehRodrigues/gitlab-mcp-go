package gitlab

import "time"

// Author
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

// User
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
	State     string `json:"state,omitempty"`
}

// Namespace
type Namespace struct {
	ID                            int    `json:"id"`
	Name                          string `json:"name"`
	Path                          string `json:"path"`
	Kind                          string `json:"kind"`
	FullPath                      string `json:"full_path"`
	ParentID                      *int   `json:"parent_id"`
	AvatarURL                     string `json:"avatar_url"`
	WebURL                        string `json:"web_url"`
	MembersCountWithDescendants   int    `json:"members_count_with_descendants,omitempty"`
	BillableMembersCount          int    `json:"billable_members_count,omitempty"`
	Plan                          string `json:"plan,omitempty"`
	Trial                         bool   `json:"trial,omitempty"`
	ProjectsCount                 int    `json:"projects_count,omitempty"`
}

type NamespaceExistsResponse struct {
	Exists  bool     `json:"exists"`
	Suggests []string `json:"suggests,omitempty"`
}

// Project / Repository
type ProjectNamespace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	AvatarURL string `json:"avatar_url,omitempty"`
	WebURL    string `json:"web_url,omitempty"`
}

type ProjectPermissions struct {
	ProjectAccess *AccessLevel `json:"project_access,omitempty"`
	GroupAccess   *AccessLevel `json:"group_access,omitempty"`
}

type AccessLevel struct {
	AccessLevel       int `json:"access_level"`
	NotificationLevel int `json:"notification_level,omitempty"`
}

type SharedGroup struct {
	GroupID          int    `json:"group_id"`
	GroupName        string `json:"group_name"`
	GroupFullPath    string `json:"group_full_path"`
	GroupAccessLevel int    `json:"group_access_level"`
}

type Project struct {
	ID                              int                  `json:"id"`
	Name                            string               `json:"name"`
	PathWithNamespace               string               `json:"path_with_namespace"`
	Visibility                      string               `json:"visibility,omitempty"`
	Owner                           *User                `json:"owner,omitempty"`
	WebURL                          string               `json:"web_url,omitempty"`
	Description                     *string              `json:"description"`
	Fork                            bool                 `json:"fork,omitempty"`
	SSHURLToRepo                    string               `json:"ssh_url_to_repo,omitempty"`
	HTTPURLToRepo                   string               `json:"http_url_to_repo,omitempty"`
	CreatedAt                       string               `json:"created_at,omitempty"`
	LastActivityAt                  string               `json:"last_activity_at,omitempty"`
	DefaultBranch                   string               `json:"default_branch,omitempty"`
	Namespace                       *ProjectNamespace    `json:"namespace,omitempty"`
	ReadmeURL                       *string              `json:"readme_url,omitempty"`
	Topics                          []string             `json:"topics,omitempty"`
	OpenIssuesCount                 int                  `json:"open_issues_count,omitempty"`
	Archived                        bool                 `json:"archived,omitempty"`
	ForksCount                      int                  `json:"forks_count,omitempty"`
	StarCount                       int                  `json:"star_count,omitempty"`
	Permissions                     *ProjectPermissions  `json:"permissions,omitempty"`
	IssuesEnabled                   bool                 `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled            bool                 `json:"merge_requests_enabled,omitempty"`
	WikiEnabled                     bool                 `json:"wiki_enabled,omitempty"`
	ContainerRegistryEnabled        bool                 `json:"container_registry_enabled,omitempty"`
	SharedWithGroups                []SharedGroup        `json:"shared_with_groups,omitempty"`
}

// Fork
type ForkParent struct {
	Name              string `json:"name"`
	PathWithNamespace string `json:"path_with_namespace"`
	Owner             *User  `json:"owner,omitempty"`
	WebURL            string `json:"web_url"`
}

type Fork struct {
	Project
	ForkedFromProject *ForkParent `json:"forked_from_project,omitempty"`
}

// File Content
type FileContent struct {
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Size            int    `json:"size"`
	Encoding        string `json:"encoding"`
	Content         string `json:"content"`
	ContentSHA256   string `json:"content_sha256"`
	Ref             string `json:"ref"`
	BlobID          string `json:"blob_id"`
	CommitID        string `json:"commit_id"`
	LastCommitID    string `json:"last_commit_id"`
	ExecuteFilemode bool   `json:"execute_filemode,omitempty"`
}

type DirectoryEntry struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	ID    string `json:"id"`
	WebURL string `json:"web_url"`
}

// CreateUpdateFileResponse
type CreateUpdateFileResponse struct {
	FilePath string       `json:"file_path"`
	Branch   string       `json:"branch"`
	CommitID string       `json:"commit_id,omitempty"`
	Content  *FileContent `json:"content,omitempty"`
}

// Tree
type TreeItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Mode string `json:"mode"`
}

type Tree struct {
	ID   string     `json:"id"`
	Tree []TreeItem `json:"tree"`
}

// Commit
type CommitStats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type CommitRef struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Commit struct {
	ID             string       `json:"id"`
	ShortID        string       `json:"short_id"`
	Title          string       `json:"title"`
	Message        string       `json:"message,omitempty"`
	AuthorName     string       `json:"author_name"`
	AuthorEmail    string       `json:"author_email"`
	AuthoredDate   string       `json:"authored_date"`
	CommitterName  string       `json:"committer_name"`
	CommitterEmail string       `json:"committer_email"`
	CommittedDate  string       `json:"committed_date"`
	WebURL         string       `json:"web_url"`
	ParentIDs      []string     `json:"parent_ids"`
	Stats          *CommitStats `json:"stats,omitempty"`
	Status         *string      `json:"status,omitempty"`
}

// Reference (Branch)
type Reference struct {
	Name   string `json:"name"`
	Commit struct {
		ID     string `json:"id"`
		WebURL string `json:"web_url"`
	} `json:"commit"`
}

// Diff
type Diff struct {
	OldPath     string `json:"old_path"`
	NewPath     string `json:"new_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	Diff        string `json:"diff"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

type CompareResult struct {
	Commit         *Commit  `json:"commit,omitempty"`
	Commits        []Commit `json:"commits"`
	Diffs          []Diff   `json:"diffs"`
	CompareTimeout bool     `json:"compare_timeout,omitempty"`
	CompareSameRef bool     `json:"compare_same_ref,omitempty"`
}

// Label
type Label struct {
	ID                    int     `json:"id"`
	Name                  string  `json:"name"`
	Color                 string  `json:"color"`
	TextColor             string  `json:"text_color"`
	Description           *string `json:"description"`
	DescriptionHTML       *string `json:"description_html,omitempty"`
	OpenIssuesCount       int     `json:"open_issues_count,omitempty"`
	ClosedIssuesCount     int     `json:"closed_issues_count,omitempty"`
	OpenMergeRequestsCount int    `json:"open_merge_requests_count,omitempty"`
	Subscribed            bool    `json:"subscribed,omitempty"`
	Priority              *int    `json:"priority,omitempty"`
	IsProjectLabel        bool    `json:"is_project_label,omitempty"`
}

// Milestone
type Milestone struct {
	ID          int        `json:"id"`
	IID         int        `json:"iid"`
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	DueDate     *string    `json:"due_date"`
	StartDate   *string    `json:"start_date"`
	State       string     `json:"state"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	Expired     bool       `json:"expired"`
	WebURL      string     `json:"web_url,omitempty"`
}

// Issue
type MilestoneSummary struct {
	ID          int     `json:"id"`
	IID         int     `json:"iid"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	State       string  `json:"state"`
	WebURL      string  `json:"web_url"`
}

type TimeStats struct {
	TimeEstimate        int     `json:"time_estimate"`
	TotalTimeSpent      int     `json:"total_time_spent"`
	HumanTimeEstimate   *string `json:"human_time_estimate"`
	HumanTotalTimeSpent *string `json:"human_total_time_spent"`
}

type Issue struct {
	ID              int               `json:"id"`
	IID             int               `json:"iid"`
	ProjectID       int               `json:"project_id"`
	Title           string            `json:"title"`
	Description     *string           `json:"description"`
	State           string            `json:"state"`
	Author          User              `json:"author"`
	Assignees       []User            `json:"assignees"`
	Labels          []any             `json:"labels"`
	Milestone       *MilestoneSummary `json:"milestone"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	ClosedAt        *time.Time        `json:"closed_at"`
	WebURL          string            `json:"web_url"`
	Confidential    bool              `json:"confidential,omitempty"`
	DueDate         *string           `json:"due_date,omitempty"`
	DiscussionLocked *bool            `json:"discussion_locked,omitempty"`
	Weight          *int              `json:"weight,omitempty"`
}

// Issue Link
type IssueLink struct {
	SourceIssue Issue  `json:"source_issue"`
	TargetIssue Issue  `json:"target_issue"`
	LinkType    string `json:"link_type"`
}

type IssueWithLinkDetails struct {
	Issue
	IssueLinkID    int    `json:"issue_link_id"`
	LinkType       string `json:"link_type"`
	LinkCreatedAt  string `json:"link_created_at"`
	LinkUpdatedAt  string `json:"link_updated_at"`
}

// Merge Request
type DiffRefs struct {
	BaseSHA  string `json:"base_sha"`
	HeadSHA  string `json:"head_sha"`
	StartSHA string `json:"start_sha"`
}

type MergeRequest struct {
	ID                            int        `json:"id"`
	IID                           int        `json:"iid"`
	ProjectID                     int        `json:"project_id"`
	Title                         string     `json:"title"`
	Description                   *string    `json:"description"`
	State                         string     `json:"state"`
	Merged                        bool       `json:"merged,omitempty"`
	Draft                         bool       `json:"draft,omitempty"`
	Author                        User       `json:"author"`
	Assignees                     []User     `json:"assignees,omitempty"`
	Reviewers                     []User     `json:"reviewers,omitempty"`
	SourceBranch                  string     `json:"source_branch"`
	TargetBranch                  string     `json:"target_branch"`
	DiffRefs                      *DiffRefs  `json:"diff_refs,omitempty"`
	WebURL                        string     `json:"web_url"`
	CreatedAt                     time.Time  `json:"created_at"`
	UpdatedAt                     time.Time  `json:"updated_at"`
	MergedAt                      *time.Time `json:"merged_at"`
	ClosedAt                      *time.Time `json:"closed_at"`
	MergeCommitSHA                *string    `json:"merge_commit_sha"`
	DetailedMergeStatus           string     `json:"detailed_merge_status,omitempty"`
	MergeStatus                  string     `json:"merge_status,omitempty"`
	MergeError                   *string    `json:"merge_error,omitempty"`
	WorkInProgress               bool       `json:"work_in_progress,omitempty"`
	BlockingDiscussionsResolved  bool       `json:"blocking_discussions_resolved,omitempty"`
	ShouldRemoveSourceBranch     *bool      `json:"should_remove_source_branch,omitempty"`
	ForceRemoveSourceBranch      *bool      `json:"force_remove_source_branch,omitempty"`
	AllowCollaboration           bool       `json:"allow_collaboration,omitempty"`
	MergeWhenPipelineSucceeds    bool       `json:"merge_when_pipeline_succeeds,omitempty"`
	Squash                       bool       `json:"squash,omitempty"`
	Labels                       []string   `json:"labels,omitempty"`
	ChangesCount                 *string    `json:"changes_count,omitempty"`
}

// Discussion / Note
type NotePosition struct {
	BaseSHA      string    `json:"base_sha"`
	StartSHA     string    `json:"start_sha"`
	HeadSHA      string    `json:"head_sha"`
	OldPath      string    `json:"old_path,omitempty"`
	NewPath      string    `json:"new_path,omitempty"`
	PositionType string    `json:"position_type"`
	OldLine      *int      `json:"old_line,omitempty"`
	NewLine      *int      `json:"new_line,omitempty"`
	Width        int       `json:"width,omitempty"`
	Height       int       `json:"height,omitempty"`
	X            int       `json:"x,omitempty"`
	Y            int       `json:"y,omitempty"`
}

type Note struct {
	ID           int           `json:"id"`
	Type         *string       `json:"type"`
	Body         string        `json:"body"`
	Attachment   any           `json:"attachment"`
	Author       User          `json:"author"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	System       bool          `json:"system"`
	NoteableID   int           `json:"noteable_id"`
	NoteableType string        `json:"noteable_type"`
	ProjectID    int           `json:"project_id,omitempty"`
	NoteableIID  *int          `json:"noteable_iid"`
	Resolvable   bool          `json:"resolvable,omitempty"`
	Resolved     bool          `json:"resolved,omitempty"`
	ResolvedBy   *User         `json:"resolved_by,omitempty"`
	ResolvedAt   *time.Time    `json:"resolved_at,omitempty"`
	Position     *NotePosition `json:"position,omitempty"`
}

type Discussion struct {
	ID             string `json:"id"`
	IndividualNote bool   `json:"individual_note"`
	Notes          []Note `json:"notes"`
}

// Pipeline
type Pipeline struct {
	ID        int     `json:"id"`
	ProjectID int     `json:"project_id"`
	SHA       string  `json:"sha"`
	Ref       string  `json:"ref"`
	Status    string  `json:"status"`
	Source    string  `json:"source,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	WebURL    string  `json:"web_url"`
	Duration  *int    `json:"duration,omitempty"`
	StartedAt *string `json:"started_at,omitempty"`
	FinishedAt *string `json:"finished_at,omitempty"`
	Coverage  *float64 `json:"coverage,omitempty"`
}

// Pipeline Job
type PipelineJob struct {
	ID        int        `json:"id"`
	Status    string     `json:"status"`
	Stage     string     `json:"stage"`
	Name      string     `json:"name"`
	Ref       string     `json:"ref"`
	Tag       bool       `json:"tag"`
	Coverage  *float64   `json:"coverage,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Duration  *float64   `json:"duration,omitempty"`
	WebURL    string     `json:"web_url,omitempty"`
}

// Wiki Page
type WikiPage struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Format    string `json:"format"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Search Response
type SearchResponse struct {
	Count       int       `json:"count"`
	TotalPages  int       `json:"total_pages"`
	CurrentPage int       `json:"current_page"`
	Items       []Project `json:"items"`
}

// Pagination
type Pagination struct {
	NextPage    *int `json:"x_next_page"`
	Page        int  `json:"x_page,omitempty"`
	PerPage     int  `json:"x_per_page,omitempty"`
	PrevPage    *int `json:"x_prev_page"`
	Total       *int `json:"x_total"`
	TotalPages  *int `json:"x_total_pages"`
}

type PaginatedDiscussions struct {
	Items      []Discussion `json:"items"`
	Pagination Pagination   `json:"pagination"`
}

// Users response
type UsersResponse map[string]*User

// Commit comment
type CommitComment struct {
	Note      string `json:"note"`
	Author    User   `json:"author"`
	CreatedAt string `json:"created_at,omitempty"`
	LineType  string `json:"line_type,omitempty"`
	Line      *int   `json:"line,omitempty"`
	Path      string `json:"path,omitempty"`
}

// File operation for push_files
type FileOperation struct {
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}
