export interface AddRepo {
  repoId: number;
  displayName: string;
}

export interface WorkflowRun {
  display_title: string;
  status: string;
  conclusion: string;
  created_at: string;
}

export interface Actions {
  total_count: number;
  workflow_runs: WorkflowRun[];
}

export interface Repo {
  id: number;
  name: string;
  actions: Actions;
}
