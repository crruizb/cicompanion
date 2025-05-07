import RepoList from "@/features/dashboard/RepoList";
import AddRepoModal from "@/features/repos/AddRepoModal";
import Cookies from "js-cookie";
import { User } from "lucide-react";
import React from "react";

const Dashboard: React.FC = () => {
  const username = Cookies.get("username");

  return (
    <main className="m-4">
      <div className="flex items-center gap-2">
        <User size={30} className="text-neutral-400" />{" "}
        <h1 className="text-xl font-extralight">{username}'s Dashboard</h1>
      </div>
      <div className="w-full h-px bg-gray-200 my-4 mb-5" />
      <header className="mb-4">
        <AddRepoModal />
      </header>
      <RepoList />
    </main>
  );
};

export default Dashboard;
