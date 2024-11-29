import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import Navbar from "@/components/ui/navbar";

test('renders greeting with name', () => {
    render(<Navbar prop={{profilePicture: String("/areaLogo.png")}}/>);
    expect(screen.getByText('Hello, John!')).toBeInTheDocument();
});
