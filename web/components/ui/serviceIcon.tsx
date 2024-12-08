import { serviceIcons } from '@/lib/serviceIcons';
import {IconType} from "react-icons";

interface ServiceIconProps {
    tag: string;
    className?: string;
}

export const ServiceIcon: React.FC<ServiceIconProps> = ({ tag, className }) => {
    const IconComponent = serviceIcons[tag.toLowerCase()] as IconType | undefined;

    return (
        IconComponent && <IconComponent className={`base-input-styles ${className || ''}`}/>
    );
};
