import { Header } from "@/features/landing/Header";
import { Hero } from "@/features/landing/Hero";
import { Features } from "@/features/landing/Features";
import { HowItWorks } from "@/features/landing/HowItWorks";
import { Footer } from "@/features/landing/Footer";

const Landing = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50">
      <Header />
      <Hero />
      <Features />
      <HowItWorks />
      <Footer />
    </div>
  );
};

export default Landing;
