import { Button } from "@/components/ui/button";
import { API_URL } from "@/services/repos";
import { Github, Activity } from "lucide-react";

export const Header = () => {
  const loginWithGithub = () => {
    window.location.href = `${API_URL}/auth/github/login`;
  };

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-white/80 backdrop-blur-sm">
      <div className="mx-6 px-1">
        <div className="flex h-16 items-center justify-between">
          <div className="flex items-center space-x-2">
            <div className="h-8 w-8 rounded-lg bg-gradient-to-br from-blue-600 to-indigo-600 flex items-center justify-center">
              <Activity className="h-5 w-5 text-white" />
            </div>
            <span className="text-xl font-bold text-gray-900">
              CI Companion
            </span>
          </div>

          <Button
            className="bg-gray-900 hover:bg-gray-600 text-white cursor-pointer"
            onClick={loginWithGithub}
          >
            <Github className="h-4 w-4 mr-2" />
            Go to your dashboard
          </Button>
        </div>
      </div>
    </header>
  );
};
