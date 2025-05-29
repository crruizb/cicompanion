import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Github, Activity, Shield, Zap, Eye, GitBranch } from "lucide-react";

export const Features = () => {
  const features = [
    {
      icon: Github,
      title: "GitHub Integration",
      description:
        "Easily connect with your GitHub account and securely access all your repositories.",
      color: "from-gray-600 to-gray-800",
    },
    {
      icon: Activity,
      title: "Real-time monitoring",
      description:
        "See the current status of all your GitHub Actions and receive instant updates.",
      color: "from-green-500 to-emerald-600",
    },
    {
      icon: Eye,
      title: "Unified view",
      description:
        "All your repositories and their CI/CD status in one organized, clear screen.",
      color: "from-blue-500 to-indigo-600",
    },
    {
      icon: GitBranch,
      title: "Run history",
      description:
        "Access complete execution history and analyze trends in your builds.",
      color: "from-purple-500 to-violet-600",
    },
    {
      icon: Zap,
      title: "Quick setup",
      description:
        "Get started in seconds. Just authorize with GitHub and select repositories.",
      color: "from-yellow-500 to-orange-500",
    },
    {
      icon: Shield,
      title: "Secure and private",
      description:
        "Your data is safe. We only access the minimum information needed.",
      color: "from-red-500 to-pink-600",
    },
  ];

  return (
    <section id="features" className="py-20 bg-white">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
            Everything you need to monitor your CI/CD
          </h2>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto">
            CICompanion simplifies tracking your GitHub Actions with powerful
            tools and an intuitive interface.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {features.map((feature, index) => (
            <Card
              key={index}
              className="group hover:shadow-lg transition-all duration-300 hover:-translate-y-1 border-0 shadow-md"
            >
              <CardHeader>
                <div
                  className={`h-12 w-12 rounded-lg bg-gradient-to-br ${feature.color} flex items-center justify-center mb-4 group-hover:scale-110 transition-transform duration-300`}
                >
                  <feature.icon className="h-6 w-6 text-white" />
                </div>
                <CardTitle className="text-xl font-semibold text-gray-900">
                  {feature.title}
                </CardTitle>
              </CardHeader>
              <CardContent>
                <CardDescription className="text-gray-600 leading-relaxed">
                  {feature.description}
                </CardDescription>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
};
