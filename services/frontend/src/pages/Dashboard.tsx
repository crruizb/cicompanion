import RepoList from "@/features/dashboard/RepoList";
import AddRepoModal from "@/features/repos/AddRepoModal";

export default function Dashboard() {
  return (
    <main className="m-4">
      <header>
        <AddRepoModal />
      </header>
      <RepoList />
    </main>
  );
}
