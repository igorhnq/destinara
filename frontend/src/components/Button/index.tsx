interface ButtonProps {
    children: React.ReactNode;
    borderColor?: string;
    textColor?: string;  
}

export default function Button({ children, borderColor = "border-white", textColor = "text-white" }: ButtonProps) {
    return (
        <button
            className={`flex items-center justify-center ${borderColor} ${textColor} text-xl font-bold border-4 rounded-3xl p-2 w-60 hover:cursor-pointer`}
        >
            {children}
        </button>
    );
}