import { EnvelopeSimpleIcon, LockSimpleIcon, UserIcon } from "@phosphor-icons/react";
import FormInput from "../../components/FormInput";
import DefaultLayout from "../../layouts/DefaultLayout";
import Button from "../../components/Button";

export default function Register() {
    return (
        <DefaultLayout>
            <div className="flex-1 flex flex-col gap-5 items-center justify-center text-center h-full">
                <div className="flex flex-col">
                    <h2 className="text-[#7DDADD] text-2xl font-bold">CRIE SUA CONTA</h2>
                    <p className="text-[#939090] text-sm">Preencha seus dados</p>
                </div>
                <form action="" className="flex flex-col gap-4">
                    <FormInput
                        icon={<UserIcon size={24} color="#939090" weight="bold" />}
                        placeholder="Nome"
                    />
                    <FormInput
                        icon={<EnvelopeSimpleIcon size={24} color="#939090" weight="bold" />}
                        placeholder="E-mail"
                    />
                    <FormInput
                        icon={<LockSimpleIcon size={24} color="#939090" weight="bold" />}
                        placeholder="Senha"
                    />
                </form>
                <Button borderColor="border-[#7DDADD]" textColor="text-[#7DDADD]">
                    CRIAR CONTA
                </Button>
            </div>
        </DefaultLayout>
    )
}