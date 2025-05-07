import "./index.css";
import Cookies from "js-cookie";
import { Button } from "@/components/ui/button";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Dashboard from "./pages/Dashboard";
import { Toaster } from "react-hot-toast";

function App() {
  const loginWithGithub = () => {
    window.location.href = "http://localhost:8080/auth/github/login";
  };

  const username = Cookies.get("username");

  const queryClient = new QueryClient();

  return (
    <>
      {!username && (
        <Button variant="outline" onClick={loginWithGithub}>
          Login with Github
        </Button>
      )}

      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <Routes>
            <Route index element={<Navigate replace to="dashboard" />} />
            <Route path="dashboard" element={<Dashboard />} />
          </Routes>
        </BrowserRouter>
        <Toaster
          position="top-center"
          gutter={12}
          containerStyle={{ margin: "8px" }}
          toastOptions={{
            success: {
              duration: 3000,
            },
            error: {
              duration: 5000,
            },
            style: {
              fontSize: "16px",
              maxWidth: "500px",
              padding: "16px 24px",
            },
          }}
        />
      </QueryClientProvider>
    </>
  );
}

export default App;
