import { UserIcon } from "@phosphor-icons/react";

interface HeaderProps {
    isPackageDetails?: boolean;
    packageName?: string;
}

export default function Header({ isPackageDetails = false, packageName }: HeaderProps) {
    return (
        <header className="flex items-center justify-between w-full">
            <div className="flex items-center gap-5 h-[75px]">
                <div className="flex items-center justify-center border-[#5BB8BB] border-2 w-[75px] h-[75px]">logo</div>
                {!isPackageDetails && (
                    <h1 className="text-[#5BB8BB] text-5xl font-bold">PACOTE DE VIAGENS</h1>
                )}
                {isPackageDetails && packageName && (
                    <div className="flex flex-col gap-2">
                        <h1 className="text-[#5BB8BB] text-5xl font-bold">{packageName}</h1>
                        <p className="text-[#939090] text-sm">Manaus é o principal ponto de entrada para a Amazônia, oferecendo uma mistura de cultura indígena e belezas naturais. </p>
                    </div>
                )}
            </div>
            <UserIcon size={42} color="#777777" weight="bold" />
        </header>
    )
}