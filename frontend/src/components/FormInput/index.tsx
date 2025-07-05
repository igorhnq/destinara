interface FormInputProps {
  icon: React.ReactNode;
  placeholder: string;
}

export default function FormInput({ icon, placeholder }: FormInputProps) {
    return (
        <div className="bg-[#DBDBDB] flex items-center gap-2 p-4">
            {icon}
            <input type="text" placeholder={placeholder} className="w-md placeholder:text-[#939090] placeholder:font-medium bg-transparent outline-none text-[#939090] font-semibold" />
        </div>
    )
}