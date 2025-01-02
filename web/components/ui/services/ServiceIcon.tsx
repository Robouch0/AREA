import { serviceIcons } from '@/lib/serviceIcons';
import { IconType } from "react-icons";
import React from "react";

interface ServiceIconProps {
    tag: string;
    className?: string;
    size?: number; // Add a size prop
}

export const ServiceIcon: React.FC<ServiceIconProps> = ({ tag, className, size = 24 }) => {
    const IconComponent = serviceIcons[tag.toLowerCase()] as IconType | undefined;

    return IconComponent ? (
        <div style={{ fontSize: `${size}px`, width: `${size}px`, height: `${size}px` }} className={className}>
            <IconComponent style={{ width: '100%', height: '100%' }} />
        </div>
    ) : null;
};
