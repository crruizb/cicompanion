let API_URL = "https://cicompanion.com/api";
const isLocalhost =
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1";
if (isLocalhost) {
  API_URL = "http://localhost:8080/api";
}

export async function getRepos() {
  const res = await fetch(`${API_URL}/repos`, {
    credentials: "include",
  });

  if (!res.ok) throw Error("Could not fetch Github repositories");

  const { repos } = await res.json();
  return repos;
}

export async function addRepo(repoId: number, displayName: string) {
  const res = await fetch(`${API_URL}/repos`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      id: repoId,
      name: displayName,
    }),
  });

  console.log(res);
  if (!res.ok) throw Error("Could not add Github repository");
}

export async function deleteRepo(repoId: number) {
  const res = await fetch(`${API_URL}/repos/${repoId}`, {
    method: "DELETE",
    credentials: "include",
  });

  console.log(res);
  if (!res.ok) throw Error("Could not add Github repository");
}

export async function getDashboard() {
  const res = await fetch(`${API_URL}/dashboard/repos`, {
    credentials: "include",
  });

  if (!res.ok) throw Error("Could not fetch user dashboard");

  const { repos } = await res.json();
  return repos;
}
