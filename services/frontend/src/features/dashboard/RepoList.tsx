import { useDashboard } from "../repos/useRepos";
import ActionsBanner from "./ActionsBanner";

export default function RepoList() {
  const { data } = useDashboard();
  return (
    <section className="grid grid-cols-2 gap-4">
      {data &&
        data.map((repo) => (
          <ActionsBanner
            key={repo.id}
            id={repo.id}
            name={repo.name}
            actions={repo.actions}
          />
        ))}
    </section>
  );
}
