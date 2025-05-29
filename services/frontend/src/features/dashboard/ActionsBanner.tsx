import React from "react";
import {
  Check,
  X,
  Calendar,
  FileText,
  AlertCircle,
  Trash2,
} from "lucide-react";
import { Card } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { format } from "date-fns";
import { useDeleteRepo } from "@/features/repos/useRepos";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import toast from "react-hot-toast";

interface WorkflowRun {
  display_title: string;
  status: string;
  conclusion: string;
  created_at: string;
}

interface Actions {
  total_count: number;
  workflow_runs: WorkflowRun[];
}

interface ActionsBannerProps {
  id: number;
  name: string;
  actions: Actions;
}

const ActionsBanner: React.FC<ActionsBannerProps> = ({ id, name, actions }) => {
  const { removeRepo } = useDeleteRepo();
  const [showDeleteDialog, setShowDeleteDialog] = React.useState(false);

  // Check if workflow_runs array is empty
  const hasWorkflowRuns =
    actions.workflow_runs && actions.workflow_runs.length > 0;
  const latestRun = hasWorkflowRuns ? actions.workflow_runs[0] : null;

  // Default to neutral color if no runs
  const isSuccess = latestRun?.conclusion === "success";
  const borderColorClass = !hasWorkflowRuns
    ? "border-gray-300"
    : isSuccess
    ? "border-green-500"
    : "border-red-500";

  // Format the date to be more readable
  const formattedDate = latestRun?.created_at
    ? format(new Date(latestRun.created_at), "MMM d, yyyy 'at' h:mm a")
    : "N/A";

  const handleSubmit = (id: number) => {
    setShowDeleteDialog(false);
    removeRepo(id);
  };

  return (
    <Card
      className={cn(
        "relative p-5 border-l-4 transition-all",
        borderColorClass,
        "hover:shadow-lg"
      )}
    >
      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div className="flex flex-col">
          <div className="flex items-center gap-2 mb-1">
            <FileText size={18} className="text-gray-600" />
            <h3 className="font-semibold text-lg">{name}</h3>
          </div>

          {hasWorkflowRuns ? (
            <>
              <div className="flex items-center gap-2 text-sm text-gray-600 mb-1">
                <FileText size={16} className="text-gray-500" />
                <span>{latestRun?.display_title}</span>
              </div>

              <div className="flex items-center gap-2 text-sm text-gray-500">
                <Calendar size={16} />
                <span>{formattedDate}</span>
              </div>
            </>
          ) : (
            <div className="flex items-center gap-2 text-sm text-gray-500 mt-1">
              <AlertCircle size={16} />
              <span>No workflow runs available yet</span>
            </div>
          )}
        </div>

        <div className="flex flex-col sm:items-end gap-2">
          {hasWorkflowRuns ? (
            <>
              <div className="flex items-center gap-2">
                {isSuccess ? (
                  <Check size={20} className="text-green-500" />
                ) : (
                  <X size={20} className="text-red-500" />
                )}
                <span
                  className={cn(
                    "font-medium",
                    isSuccess ? "text-green-500" : "text-red-500"
                  )}
                >
                  {latestRun?.conclusion.toUpperCase()}
                </span>
              </div>
              <div className="text-sm text-gray-500">
                Total runs: {actions.total_count}
              </div>
            </>
          ) : (
            <div className="text-sm text-gray-500">Total runs: 0</div>
          )}
        </div>
      </div>

      <button
        onClick={() => setShowDeleteDialog(true)}
        className="absolute top-2 right-2 p-1 rounded-full text-gray-400 hover:text-red-500 hover:bg-gray-100 transition-colors cursor-pointer"
        aria-label="Delete repository"
      >
        <Trash2 size={16} />
      </button>

      <Dialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Repository</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete "{name}"?
            </DialogDescription>
          </DialogHeader>
          <div className="flex justify-end gap-3 mt-4">
            <DialogClose asChild>
              <Button variant="outline" className="cursor-pointer">
                Cancel
              </Button>
            </DialogClose>
            <Button
              variant="destructive"
              className="cursor-pointer"
              onClick={() => handleSubmit(id)}
            >
              Delete
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </Card>
  );
};

export default ActionsBanner;
