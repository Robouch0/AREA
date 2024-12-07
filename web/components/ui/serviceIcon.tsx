import { serviceIcons } from '@/lib/serviceIcons';
import {IconType} from "react-icons";

interface ServiceIconProps {
    tag: string;
}

export const ServiceIcon: React.FC<ServiceIconProps> = ({ tag }) => {
    const IconComponent = serviceIcons[tag.toLowerCase()] as IconType | undefined;

    return (
        <>
            {IconComponent && <IconComponent className="text-black text-2xl" />}
        </>
    );
};
