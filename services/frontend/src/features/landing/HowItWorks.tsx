import { Card, CardContent } from "@/components/ui/card";
import { Github, Plus, Eye } from "lucide-react";

export const HowItWorks = () => {
  const steps = [
    {
      step: "1",
      icon: Github,
      title: "Sign in with GitHub",
      description:
        "Securely connect your GitHub account with a single click. We only need read permissions for your repositories.",
      color: "from-blue-500 to-indigo-600",
    },
    {
      step: "2",
      icon: Plus,
      title: "Add repositories",
      description:
        "Select the repositories you want to monitor. Add as many as you need and organize them however you prefer.",
      color: "from-green-500 to-emerald-600",
    },
    {
      step: "3",
      icon: Eye,
      title: "Monitor your Actions",
      description:
        "View the status of all GitHub Actions runs in real-time. Get notifications when something needs your attention.",
      color: "from-purple-500 to-violet-600",
    },
  ];

  return (
    <section
      id="how-it-works"
      className="py-20 bg-gradient-to-br from-gray-50 to-blue-50"
    >
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-16">
          <h2 className="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
            How it works
          </h2>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto">
            In just three simple steps you'll have complete control over your
            GitHub Actions.
          </p>
        </div>

        <div className="max-w-4xl mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 relative">
            {/* Connecting lines */}
            <div className="hidden md:block absolute top-24 left-1/6 right-1/6 h-0.5 bg-gradient-to-r from-blue-200 via-green-200 to-purple-200"></div>

            {steps.map((step, index) => (
              <div key={index} className="relative">
                <Card className="text-center p-6 hover:shadow-xl transition-all duration-300 hover:-translate-y-2 border-0 shadow-lg bg-white/80 backdrop-blur-sm">
                  <CardContent className="pt-6">
                    <div
                      className={`h-16 w-16 rounded-full bg-gradient-to-br ${step.color} flex items-center justify-center mx-auto mb-6 relative`}
                    >
                      <step.icon className="h-8 w-8 text-white" />
                      <div className="absolute -top-2 -right-2 h-8 w-8 bg-white rounded-full flex items-center justify-center text-sm font-bold text-gray-900 shadow-md">
                        {step.step}
                      </div>
                    </div>
                    <h3 className="text-xl font-semibold text-gray-900 mb-3">
                      {step.title}
                    </h3>
                    <p className="text-gray-600 leading-relaxed">
                      {step.description}
                    </p>
                  </CardContent>
                </Card>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};
