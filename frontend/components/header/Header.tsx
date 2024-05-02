"use client";

import Link from "next/link";
import { Button } from "../ui/button";

const Header = () => {
  return (
    <header className="flex items-center justify-between py-4 px-36 border-b-2 border-slate-300 bg-slate-400 min-w-full fixed top-0 left-0 right-0">
      <Link href="/">
        {" "}
        <h1 className="text-3xl font-bold px-4">Fullstack Auth Demo</h1>
      </Link>
      <Link href="/login">
        <Button variant="default">Login</Button>
      </Link>
    </header>
  );
};

export default Header;
