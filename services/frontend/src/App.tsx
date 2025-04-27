import "./index.css";
import Cookies from "js-cookie";
import { Button } from "@/components/ui/button";

function App() {
  const loginWithGithub = () => {
    window.location.href = "http://localhost:8080/auth/github/login";
  };

  const username = Cookies.get("username");

  return (
    <>
      {!username && (
        <Button variant="outline" onClick={loginWithGithub}>
          Login with Github
        </Button>
      )}
      {username && <h1>Welcome {username}</h1>}
    </>
  );
}

export default App;
