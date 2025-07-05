interface ButtonProps {
    children: React.ReactNode;
}

export default function Button({ children }: ButtonProps) {
    return (
        <button className="flex items-center justify-center border-white text-white text-xl font-bold border-4 rounded-3xl p-2 w-60 hover:cursor-pointer">
            {children}
        </button>
    );
}