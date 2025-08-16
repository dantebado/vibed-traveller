import React from 'react';

function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-gray-800 mb-4">
          Hello World!
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          Welcome to Vibed Traveller
        </p>
        <div className="bg-white rounded-lg shadow-lg p-8 max-w-md mx-auto">
          <div className="text-4xl mb-4">✈️</div>
          <p className="text-gray-700">
            Your journey begins here. This is a simple React app with TypeScript and Tailwind CSS.
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
