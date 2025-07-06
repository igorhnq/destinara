import Button from "../Button";

interface CardProps {
    name: string;
    image?: string;
}

export default function Card({ name, image }: CardProps) {
    return (
        <div 
            className="rounded-4xl overflow-hidden flex flex-col items-center justify-between p-10 text-center w-1/4 h-full bg-cover bg-center bg-no-repeat"
            style={image ? { backgroundImage: `url(${image})` } : {}}
        >
            <h1 className="text-white text-2xl font-bold">{name}</h1>
            <Button textColor="text-white" borderColor="border-white" width="w-50">
                VER MAIS
            </Button>
        </div>
    );
}