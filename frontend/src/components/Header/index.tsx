import { UserIcon } from "@phosphor-icons/react";

export default function Header() {
    return (
        <header className="flex items-center justify-between w-full">
            <div className="flex items-center gap-5">
                <div className="flex items-center justify-center border-[#5BB8BB] border-2 w-[75px] h-[75px]">logo</div>
                <h1 className="text-[#5BB8BB] text-5xl font-bold">PACOTE DE VIAGENS</h1>
            </div>
            <UserIcon size={42} color="#777777" weight="bold" />
        </header>
    )
}