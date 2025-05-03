import { useDashboard } from "../repos/useRepos";
import Repo from "./Repo";

export default function RepoList() {
  const { data } = useDashboard();
  return (
    <section className="grid grid-cols-4">
      {data &&
        data.map((repo) => (
          <Repo key={repo.id} name={repo.name} actions={repo.actions} />
        ))}
    </section>
  );
}
