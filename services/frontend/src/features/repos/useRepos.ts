import { addRepo, deleteRepo, getDashboard, getRepos } from "@/services/repos";
import { AddRepo } from "@/types/types";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

export function useRepos() {
  const { data, error } = useQuery({
    queryKey: ["repos"],
    queryFn: () => getRepos(),
  });

  return { data, error };
}

export function useAddRepo() {
  const queryClient = useQueryClient();
  const { mutate: createRepo } = useMutation({
    mutationFn: ({ repoId, displayName }: AddRepo) =>
      addRepo(repoId, displayName),
    onSuccess: () => {
      console.log("Repo added!");
      toast.success("Repo added successfully!");
      queryClient.invalidateQueries({ queryKey: ["dashboard"] });
    },
    onError: () => {
      toast.error("Could not add the repo. Try again later.");
    },
  });

  return { createRepo };
}

export function useDeleteRepo() {
  const queryClient = useQueryClient();
  const { mutate: removeRepo } = useMutation({
    mutationFn: (repoId: number) => deleteRepo(repoId),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["dashboard"] });
      toast.success("Repo removed successfully!");
    },
    onError: () => {
      toast.error("Could not remove the repo. Try again later.");
    },
  });

  return { removeRepo };
}

export function useDashboard() {
  const { data, error, isFetching } = useQuery({
    queryKey: ["dashboard"],
    queryFn: () => getDashboard(),
  });

  return { data, error, isFetching };
}
