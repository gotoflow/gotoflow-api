CREATE TABLE dag (
	DagId varchar PRIMARY KEY,
	Path varchar,
	Description varchar,
	Schedule varchar,
	StartDate Time,
	EndDate Time,
	Concurrency int,
	MaxActiveTasks int,
	TimeOut int,
	Catchup bool,
	Tags text[],
	Nodes jsonb,
    deleted_at time
)