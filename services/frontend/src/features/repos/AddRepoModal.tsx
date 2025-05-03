import { Button } from "@/components/ui/button";
import {
  DialogHeader,
  DialogFooter,
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogTitle,
} from "@/components/ui/dialog";
import { useAddRepo, useRepos } from "./useRepos";
import RepoSearch from "./RepoSearch";
import React from "react";

export default function AddRepoModal() {
  const [repo, setRepo] = React.useState<{
    displayName: string;
    repoId: number;
  }>({
    displayName: "",
    repoId: -1,
  });
  const [open, setOpen] = React.useState(false);

  const { data } = useRepos();
  const { createRepo } = useAddRepo();

  const handleRepoChange = (repoId: number, repoName: string) => {
    setRepo({ repoId: repoId, displayName: repoName });
  };

  const handleSubmit = () => {
    setOpen(false);
    createRepo(repo);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">Add repo</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add new repo</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <RepoSearch
            repos={data}
            repoName={repo.displayName}
            onRepoChange={handleRepoChange}
          />
        </div>
        <DialogFooter>
          <Button
            type="submit"
            onClick={handleSubmit}
            disabled={!repo.displayName}
          >
            Save changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
