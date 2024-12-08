import { FormFieldProps } from "@/lib/type";
import {Input} from "@/components/ui/input";

const FormField: React.FC<FormFieldProps> = ({
     type,
     placeholder,
     name,
     register,
     error,
     valueAsNumber,
     ariaLabel,
     className,
 }) => (
    <>
        <Input
            type={type}
            placeholder={placeholder}
            aria-label={ariaLabel}
            {...register(name, { valueAsNumber })}
            className={`base-input-styles ${className || ''}`}
        />
        {error && <span className="error-message text-red-500 mx-2 font-bold">{error.message}</span>}
    </>
);
export default FormField;
