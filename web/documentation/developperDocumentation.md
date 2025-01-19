# Frontend Documentation

This document provides an overview of the project structure and explains where to find key components and functionality.

## Project Structure

### Root Directory Files
- `next.config.ts` - Next.js configuration file
- `tailwind.config.ts` - Tailwind CSS configuration
- `package.json` - Project dependencies and scripts
- `tsconfig.json` - TypeScript configuration
- `postcss.config.mjs` - PostCSS configuration for Tailwind
- `components.json` - Shadcn/UI components configuration

### Core Directories

#### `/app`
The main application directory using Next.js 13+ app router structure:
- `/api` - API route handlers for form submission and token management
- `/fonts` - Custom font files (Geist)
- `/services` - Service-related pages including:
  - `/contact` - Contact page
  - `/create` - Service creation interface
  - `/faq` - FAQ page
  - `/myareas` - User's areas dashboard
  - `/profile` - User profile management
- `/register` - User registration page
- `/loadOauth` - OAuth authentication handling
- `layout.tsx` - Root layout component
- `globals.css` - Global styles

#### `/components`
Component library organized by functionality:

**Pages Components** (`/components/pages`):
- `/create` - Components for service creation
  - `CalendarTimeInput.tsx` - Date/time input component
  - `ComboboxDemo.tsx` - Searchable dropdown component
  - `CreatePage.tsx` - Main creation page component
  - Other input-related components
- `/login` - Authentication components
- `/myareas` - Area management components
- `/profile` - User profile components
- `/register` - Registration form components

**UI Components** (`/components/ui`):
- `/layouts` - Layout components including navigation (navBar)
- `/services`
  - `/areaCards` - Various card components for displaying services
  - `/oauth` - OAuth-related components
- `/utils` - Utility components including:
  - Form fields
  - Service icons
  - Time picker components
  - `/thirdPartyComponents` - Third-party component implementations
    - `/shadcn` - All the imported shadcn ui components


#### `/api`
API integration layer:
- Authentication handlers
- Area management (create, delete, enable)
- User information management
- Type definitions for API requests/responses

### Development Files
- `/public` - Static assets
- `/tests` - Test files
- `/hooks` - Custom React hooks
- `/lib` - Shared utilities and libraries

## Getting Started

1. Install dependencies:
```bash
npm install
```

2. Run the development server:
```bash
npm run dev
```

3. Build for production:
```bash
npm run build
```

## Key Features

- **Authentication System**: Implements OAuth authentication flow
- **Service Management**: Create, manage, and monitor service areas
- **Profile Management**: User profile customization
- **Responsive UI**: Built with Tailwind CSS and Shadcn/UI components
- **Type Safety**: Full TypeScript implementation

## Component Guidelines

- All reusable UI components should be placed in `/components/ui`
- Page-specific components go in `/components/pages/{page-name}`
- Shared utilities should be placed in `/lib`
- If you want to include external components from another lib you just need to create a new directory with the name of the subsequent librairies in the `/components/ui/utils/thirdPartyComponents` directory
- API integration code belongs in `/api`

## Tech Stack

- Next.js 13+
- TypeScript
- Tailwind CSS
- Shadcn/UI
- React

## Best Practices

1. Follow the established directory structure
2. Use TypeScript for all new components
3. Implement responsive design using Tailwind classes
4. Maintain component modularity
5. Follow the existing naming conventions
