interface ButtonProps {
    children: React.ReactNode;
    borderColor?: string;
    textColor?: string;
    width?: string;
    backgroundColor?: string;
}

export default function Button({
    children,
    borderColor = "border-white",
    textColor = "text-white",
    width = "w-60",
    backgroundColor = ""
}: ButtonProps) {
    return (
        <button
            className={`flex items-center justify-center ${borderColor} ${textColor} text-xl font-bold border-4 rounded-3xl p-2 ${width} ${backgroundColor} hover:cursor-pointer`}
        >
            {children}
        </button>
    );
}