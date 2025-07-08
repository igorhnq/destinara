import type { ReactNode } from "react";

interface DefaultLayoutProps {
    children: ReactNode;
    flexCol?: boolean;
    hasPadding?: boolean;
}

export default function DefaultLayout({ children, flexCol = false, hasPadding = false }: DefaultLayoutProps) {
    return (
        <div className="min-h-screen bg-[#DBDBDB] flex items-center justify-center p-10">
            <div
                className={`bg-white rounded-2xl shadow-lg flex overflow-hidden w-full h-[calc(100vh-5rem)] ${flexCol ? "flex-col" : ""} ${hasPadding ? "p-16" : ""}`}
            >
                {children}
            </div>
        </div>
    )
}