import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Github } from "lucide-react";
import { API_URL } from "@/services/repos";

export const Hero = () => {
  const loginWithGithub = () => {
    window.location.href = `${API_URL}/auth/github/login`;
  };

  return (
    <section className="py-20 lg:py-32">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center max-w-4xl mx-auto">
          <Badge
            variant="secondary"
            className="mb-4 bg-blue-100 text-blue-700 border-blue-200"
          >
            ✨ Smart CI/CD monitoring
          </Badge>

          <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 mb-6 leading-tight">
            Stay in control of your
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600">
              {" "}
              GitHub Actions
            </span>
          </h1>

          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto leading-relaxed">
            CICompanion lets you easily monitor the status of your repositories
            and view the latest GitHub Actions runs in one place.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
            <Button
              size="lg"
              className="bg-gray-900 hover:bg-gray-600 text-white px-8 py-3 text-lg cursor-pointer"
              onClick={loginWithGithub}
            >
              <Github className="h-5 w-5 mr-2" />
              Get started free
            </Button>
          </div>
        </div>
      </div>
    </section>
  );
};
