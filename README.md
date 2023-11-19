# SubMan

## Description

Introducing an advanced, open-source subscription management application meticulously crafted for seamless integration into any React website. This optimized solution empowers users to effortlessly create personalized subscription plans, free of charge, with a rich set of customizable features.

Users have the flexibility to define a custom title, description, billing cycle, and price for their subscription plans. Additionally, they can set a desired service fee to be paid to operators upon successful payment processing. The service fee is fully adjustable, allowing owners to optimize profitability by exceeding gas costs.

Our innovative plugin enhances user experience by providing a modal interface on click. This modal presents a user-friendly form for entering essential inputs and subscription plan details. The entire process is designed to be visually appealing and easily adaptable with customizable color schemes, ensuring seamless integration into any React website.

Subscribing is made incredibly straightforward. Upon completing the necessary token approval, subscribers simply need to sign a typed data confirming their agreed-upon conditions. From there, operators take charge of handling the subscription throughout each cycle until the agreed duration concludes.

Our system prioritizes user autonomy, allowing subscribers to cancel their subscriptions at any time without incurring additional charges. This forward-thinking, tech-savvy solution combines user-friendliness with advanced customization, making subscription management a breeze for both owners and subscribers alike.

## How It`s Made

Introducing our subscription management system, where we leverage typed data signatures for seamless agreements and authentication, including support for smart account signatures for scalability.

Built on Golang, our backend features a concurrent indexer for Base Goerli and Scroll Sepolia chains, providing API endpoints for subscription requests and executable subscription information for operators.

The operator script, developed in TypeScript, is easily configurable for widespread usability. It efficiently fetches and executes subscription requests.

Our frontend relies on ENS Thorin's UI components, covering every aspect of the dApp. Nouns DAO visuals and fonts give a unique touch, featuring a cat logo.

Web3Modal V3 manages wallet connectivity, and our npm-packaged plugin, designed with ENS Thorin, is ready for use, needing wagmi.sh for compatibility.

We explored Certora's formal verification app during the hackathon, making strides in understanding its fundamentals and implementing a simple test in the repository.