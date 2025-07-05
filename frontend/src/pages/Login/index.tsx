import Button from "../../components/Button";
import { EnvelopeSimpleIcon, LockSimpleIcon } from "@phosphor-icons/react";
import FormInput from "../../components/FormInput";
import DefaultLayout from "../../layouts/DefaultLayout";

export default function Login() {
    return (
        <DefaultLayout>
            <div className="bg-[#7DDADD] w-[400px] h-full flex flex-col gap-7 items-center justify-center text-center p-4">
                <div className="flex items-center justify-center border-white border-2 w-[275px] h-[275px] ">logo</div>
                <h1 className="text-white text-4xl font-extrabold">Bem-vindo de volta!</h1>
                <p className="text-white text-xl font-normal">Acesse sua conta agora mesmo</p>
            </div>
            <div className="flex-1 flex flex-col gap-5 items-center justify-center text-center h-full">
                <div className="flex flex-col">
                    <h2 className="text-[#7DDADD] text-2xl font-bold">FAÇA SEU LOGIN</h2>
                    <p className="text-[#939090] text-sm">Preencha seus dados</p>
                </div>
                <form action="" className="flex flex-col gap-4">
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
                    ENTRAR
                </Button>
            </div>
        </DefaultLayout>
    )
}