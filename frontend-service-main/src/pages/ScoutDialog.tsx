// DialogForm.tsx
import React, { useState } from 'react';

interface DialogProps {
  isOpen: boolean;
  headerTitle: string;
  onClose: () => void;
  onSubmit: (formData: { github_url: string, docker_base_url: string }) => void;
}

const ScoutDialog: React.FC<DialogProps> = ({ isOpen, onClose, headerTitle, onSubmit }) => {
  const [field1, setField1] = useState('');
  const [field2, setField2] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({ github_url: field1, docker_base_url: field2 });
    onClose(); // Close the dialog after submission
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-75 z-50">
      <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
        <h2 className="text-xl font-semibold mb-4">{headerTitle}</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="field1">
              GitHub URL
            </label>
            <input
              type="text"
              id="field1"
              value={field1}
              onChange={(e) => setField1(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="field2">
              Docker URL
            </label>
            <input
              type="text"
              id="field2"
              value={field2}
              onChange={(e) => setField2(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="flex justify-end">
            <button
              type="button"
              onClick={onClose}
              className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded mr-2"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ScoutDialog;
