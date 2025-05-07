import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@radix-ui/react-popover";
import {
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  Command,
} from "cmdk";
import { ChevronsUpDown, Check } from "lucide-react";
import React from "react";

interface Repo {
  id: number;
  name: string;
}

interface Props {
  repos: Repo[];
  repoName: string;
  onRepoChange: (repoId: number, repoName: string) => void;
}

export default function RepoSearch({ repos, repoName, onRepoChange }: Props) {
  const [open, setOpen] = React.useState(false);

  const handleRepoClick = (selectedName: string) => {
    const selectedRepo = repos.find((repo) => repo.name === selectedName);
    const isSame = selectedName === repoName;

    const newName = isSame ? "" : selectedName;
    const newId = isSame ? -1 : selectedRepo?.id ?? -1;

    onRepoChange(newId, newName);
  };

  return (
    <Popover open={open} onOpenChange={setOpen} modal={true}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-xs justify-between "
        >
          {repoName
            ? repos.find((repo) => repo.name === repoName)?.name
            : "Search repo..."}
          <ChevronsUpDown className="opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-xs p-2 mx-5 border-1 rounded-md bg-white cursor-pointer z-50">
        <Command>
          <CommandInput placeholder="Search repo..." className="h-9" />
          <CommandList className="overflow-y-auto h-96">
            <CommandEmpty>No repository found.</CommandEmpty>
            <CommandGroup>
              {repos.map((repo) => (
                <CommandItem
                  key={repo.name}
                  value={repo.name}
                  onSelect={(currentValue) => {
                    handleRepoClick(
                      currentValue === repoName ? "" : currentValue
                    );
                    setOpen(false);
                  }}
                >
                  {repo.name}
                  <Check
                    className={cn(
                      "ml-auto",
                      repoName === repo.name ? "opacity-100" : "opacity-0"
                    )}
                  />
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
