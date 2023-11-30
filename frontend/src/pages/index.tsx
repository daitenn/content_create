import { Login } from "@/components/Login";
import { useEffect } from "react";
import axios from "axios";
import { CsrfToken } from "./api/model/content_model";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

export default function Home() {
  const queryClient = new QueryClient();

  useEffect(() => {
    axios.defaults.withCredentials = true;
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.NEXT_PUBLIC_API_URL}/csrf`
      );
      axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
    };
    getCsrfToken();
  }, []);
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <Login></Login>
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </>
  );
}
