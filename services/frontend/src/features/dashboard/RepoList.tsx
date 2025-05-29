import { useDashboard } from "../repos/useRepos";
import ActionsBanner from "./ActionsBanner";
import SpinnerLoader from "@/components/ui/SpinnerLoader";

export default function RepoList() {
  const { data, isFetching } = useDashboard();
  return (
    <>
      {isFetching ? <SpinnerLoader /> : ""}
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
    </>
  );
}
