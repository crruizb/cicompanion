interface WorkflowRuns {
  display_title: string;
  status: string;
  conclusion: string;
  created_at: string;
}

interface Actions {
  total_count: number;
  workflow_runs: WorkflowRuns[];
}

interface Props {
  name: string;
  actions: Actions;
}

export default function Repo({ name, actions }: Props) {
  console.log(actions);
  //   console.log(actions.workflow_runs[0].display_title);
  return (
    <article className="m-8 bg-amber-200">
      <h1>{name}</h1>
      <p>{actions.total_count}</p>
      {actions.total_count > 0 && (
        <div>
          <h2>{actions.workflow_runs[0].display_title}</h2>
          <p>
            Status: {actions.workflow_runs[0].status}, conclusion:{" "}
            {actions.workflow_runs[0].conclusion}
          </p>
          <p>Created at: {actions.workflow_runs[0].created_at}</p>
        </div>
      )}
    </article>
  );
}
