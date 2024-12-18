import { serviceIcons } from '@/lib/serviceIcons';
import {IconType} from "react-icons";
import React from "react";

interface ServiceIconProps {
    tag: string;
    className?: string;
}

export const ServiceIcon: React.FC<ServiceIconProps> = ({ tag, className } : ServiceIconProps) : React.JSX.Element|undefined => {
    const IconComponent = serviceIcons[tag.toLowerCase()] as IconType | undefined;

    return (
        IconComponent && <IconComponent className={`base-input-styles ${className || ''}`}/>
    );
};
