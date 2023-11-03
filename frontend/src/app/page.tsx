"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <>
      <div className="bg-white h-screen flex mx-auto w-full justify-center">
        <div className="my-96">
          <button className="" onClick={() => router.push("/login")}>
            aaaaaaaa
          </button>
        </div>
      </div>
    </>
  );
}
