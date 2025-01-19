# AREA

This project is part of the *Application Development module* at Epitech in third year.
His main purpose is to introduce us to the most used languages, ecosystems and tools in today's industry.
The other main goal of the project is to make us understand the importance of organization and work methodology, because we will assume the role of a Software Architect team.


In order do to this project we need to implement a software suite that functions similar to [**IFTT**](https://fr.wikipedia.org/wiki/IFTTT) and/or [**Zapier**](https://en.wikipedia.org/wiki/Zapier).


## Table of Contents
- [Contribution guidelines](#contribution-guidelines)
- [Comparative studies](#benchmarks--comparative-studies)
- [Authors](#authors)
- [Frontend web README](web/README.md)



## Contribution guidelines

Feel free to try and contribute to our repository, in order to do this you need to comply with some guidelines.
We are following the [conventionnal commits norm](https://www.conventionalcommits.org/en/v1.0.0/).

We also have github actions checking our code, for example a linter to not forgot about typing in typescript, for the frontend development part in Next.js.

If you want to check the same thing as our Github Action for the frontend part before pushing, you may run those commands in the **web/** directory.



```bash
npm ci --legacy-peer-deps
```
*Stand for **npm clean & install** permit to clean the node_modules directory and install the dependecies necessary to run this project.*
   
  


```bash
npm run build --if-present
```
*Execute our build script present in the package.json, here it's the *next build* used to build our client web, apply a linter and check for validity types, and generates statics pages, to check if our client compiles.*


Here is back-end command to run the server, you may run this command in the **server/** directory.

```bash
go run .
```


You can also suggest us new features or microservices to add to our project, our contact are down [here](#authors).


# Benchmarks / Comparative studies

## Frontend Framework Comparison

For the frontend, we compared Svelte, React, and Next.js based on the following criteria:

| Criteria                | Svelte              | React               | Next.js             |
|-------------------------|---------------------|---------------------|---------------------|
| Performance             | Excellent           | Good                | Very Good           |
| Bundle Size             | Very Small          | Large               | Medium              |
| Community               | Small               | Very Large          | Large               |
| Learning Curve          | Easy                | Moderate            | Moderate            |
| SSR                     | Yes                 | Possible            | Integrated          |
| SEO                     | Good                | Average             | Excellent           |
| Ecosystem               | Limited             | Vast                | Vast                |
| Age                     | 2016                | 2013                | 2016                |

### Final Choice: Next.js

Next.js was selected for the following reasons:

1. **React-Based**: It builds on React, providing optimized features.
2. **Integrated SSR**: Server-Side Rendering enhances performance and SEO.
3. **Large Community**: The React community facilitates support and resources.
4. **Rich Ecosystem**: A mature ecosystem supports various development needs.
5. **Team Experience**: Prior experience with React within the team.


Although Svelte offers excellent performance and a small bundle size, its smaller community and relative youth make it less suitable for beginners. Next.js allows us to leverage React's advantages while providing crucial optimizations for modern web application development.


For UI components, we will primarily use Shadcn due to its extensive customization options and professional usage. While we have experience with DaisyUI, Shadcn offers a more advanced component library. We also considered NextUI for its modern animations, but its customization limitations and less accessible code made Shadcn the preferred choice.

We will also use Tailwind CSS as an helper to facilitate styling of components.



## Backend Framework Comparison

For the backend, we compared Django, NestJS, and Go based on the following criteria:

| Criteria                | Django              | NestJS              | Go                  |
|-------------------------|---------------------|---------------------|---------------------|
| Language                | Python              | TypeScript (Node.js)| Go                |
| Architecture            | MVC                 | Modular, Dependency Injection | Simple with Goroutines |
| Performance             | Good for standard apps | Solid              | Excellent, high concurrency |
| Learning Curve          | Easy (Python, MVC)  | Complex Setup       | Simple for C/Python |
| Scalability             | Good                | Good                | Excellent            |
| Security                | Built-in protections | Manual configuration | Basic                |
| Primary Use             | Web apps, CRUD      | APIs, microservices | High-performance backend |
| Community Support       | Large community     | Good documentation   | Growing community    |

### Final Choice: Go

We chose Go for the following reasons:


1. **Performance**: Go excels in high-concurrency scenarios, making it ideal for performance-critical applications.
2. **Simplicity**: The language's straightforward syntax and structure facilitate rapid development.
3. **Scalability**: Go's goroutines allow for efficient handling of multiple tasks, making it highly scalable.
4. **Growing Community**: The increasing popularity of Go ensures a wealth of resources and community support.
5. **Learning curve**: Since we all done a lot of C in the team, and certain aspects of Go can be easily apprehended with a C background.

While Django and NestJS are robust frameworks, Go's performance and simplicity make it the best fit for our backend needs.



## Database Comparison

We evaluated MongoDB, PostgreSQL, and MariaDB based on the following criteria:

| Criteria                | MongoDB            | PostgreSQL         | MariaDB            |
|-------------------------|---------------------|---------------------|---------------------|
| Data Model              | JSON, non/semi-structured data | Structured relational tables | Tables + JSON support |
| Flexibility             | Schema-less         | Fixed schema        | Flexible schema     |
| Language                | MQL                 | SQL                 | SQL                 |
| Performance             | Strong in writes    | Strong in reads     | Balanced            |
| Indexing                | Multi-type          | Advanced            | Standard            |
| Scalability             | Horizontal          | Vertical            | Vertical            |
| Learning Curve          | Easy (JSON)         | Complex             | Simple (MySQL)      |
| Age                     | 2009                | 1995                | 2009                |
| Storage                 | Significant         | Compact             | Compact             |

### Final Choice: PostgreSQL


PostgreSQL was selected for the following reasons:


1. **ACID Transactions**: It supports ACID transactions, which are essential for data integrity.
2. **Complex Queries**: PostgreSQL excels in handling complex queries and analytics.
3. **Strong Consistency**: It guarantees strong consistency, which is crucial for our project.
4. **Community and Tools**: A large community and numerous tools support development and optimization.
5. **Team Experience**: Prior experience with PostgreSQL within the team enhances our efficiency.

PostgreSQL's strengths in data integrity and complex queries make it the ideal choice for our project's database needs.

## Mobile Development Comparison


For mobile development, we compared Kotlin and Flutter based on the following criteria:


| Criteria                | Kotlin              | Flutter             |
|-------------------------|---------------------|---------------------|
| Performance             | Very Good           | Excellent           |
| Programming Language     | Kotlin              | Dart                |
| Ecosystem/Community     | Well-developed       | Rapidly growing     |
| Cross-Platform          | Android only        | Multi-Platform (Android, iOS, Windows, macOS, Linux) |
| Learning Curve          | Easy                | Moderate            |
| Age                     | 2011                | 2017                |
| Hot Reload              | No                  | Yes                 |
| UI/UX Development       | Based on Android SDK| Native with Flutter widgets |

### Final Choice: Flutter

Flutter was chosen for the following reasons:

1. **Performance**: Flutter provides excellent performance, surpassing that of Kotlin.
2. **User Experience**: Its native widgets facilitate the creation of intuitive and engaging interfaces.
3. **Developper Experience**: The *hot reload* feature allows for real-time code changes, streamlining the development process.

While Kotlin is a strong option for native Android development, Flutter's cross-platform capabilities and superior performance make it the best choice for our mobile application.

## Launch Instructions

To launch our application, simply use Docker Compose. This will start all the necessary services (frontend, backend, and database) in the correct order with proper configurations:
```bash
docker compose up
```
This command will build and start all containers defined in our docker-compose.yml file. Once the containers are up and running, you can access the application through your web browser.

---


Our technology choices were driven by a combination of performance, ease of use, community support, and prior experience. Each selected technology aligns with our project goals, in hope that it will ensure a robust and efficient development process.

# See Also

[Backend Service Documentation](./server/docs/Developer%20Documentation.md)
[Mobile Application Documentation](./mobile/documentation/devDocumentation.md)
[Frontend Service Documentation](./web/documentation/developperDocumentation.md)
[Usage of our application](./web/documentation/userDocumentation.md)

# Authors

axel.tacheau@epitech.eu

rahul.chander@epitech.eu

hugo.duchemin@epitech.eu

matthieu.godme@epitech.eu
