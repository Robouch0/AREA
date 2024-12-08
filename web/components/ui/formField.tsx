import { FormFieldProps } from "@/lib/typeLogin";
import {Input} from "@/components/ui/input";
import {Path} from "react-hook-form";

function FormField<T extends Record<string, any>>({
  type,
  placeholder,
  name,
  register,
  error,
  valueAsNumber,
  ariaLabel,
  className,
}: FormFieldProps<T>) {
    return (
    <>
        <Input
            type={type}
            placeholder={placeholder}
            aria-label={ariaLabel}
            {...register(name as Path<T>, { valueAsNumber })}
            className={`base-input-styles ${className || ''}`}
        />
        {error && <span className="error-message text-red-500 mx-2 font-bold">{error.message}</span>}
    </>
)}
export default FormField;
