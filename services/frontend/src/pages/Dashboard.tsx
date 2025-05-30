import RepoList from "@/features/dashboard/RepoList";
import AddRepoModal from "@/features/repos/AddRepoModal";
import Cookies from "js-cookie";
import { User, Activity } from "lucide-react";
import React, { useEffect } from "react";
import { useNavigate } from "react-router";

const Dashboard: React.FC = () => {
  const username = Cookies.get("ci_username");

  const navigate = useNavigate();

  useEffect(() => {
    if (!username) {
      navigate("/");
    }
  }, [navigate, username]);

  return (
    <main className="m-4">
      <header className="sticky top-0 z-50 w-full border-b bg-white/80 backdrop-blur-sm mb-4">
        <div className=" px-1">
          <div className="flex h-16 items-center justify-between">
            <div className="flex items-center space-x-2">
              <div className="h-8 w-8 rounded-lg bg-gradient-to-br from-blue-600 to-indigo-600 flex items-center justify-center">
                <Activity className="h-5 w-5 text-white" />
              </div>
              <span className="text-xl font-bold text-gray-900">
                CI Companion
              </span>
            </div>
            <div className="flex items-center gap-2">
              <User size={30} className="text-neutral-400" />{" "}
              <h1 className="text-xl font-extralight">
                {username}'s Dashboard
              </h1>
            </div>
          </div>
        </div>
      </header>

      <section className="mb-4">
        <AddRepoModal />
      </section>
      <RepoList />
    </main>
  );
};

export default Dashboard;
