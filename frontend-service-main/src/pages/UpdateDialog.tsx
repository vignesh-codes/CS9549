import React, { useState } from 'react';

interface DialogProps {
  isOpen: boolean;
  name: string;
  image: string;
  onClose: () => void;
  onSubmit: (formData: { name: string; replicas: number; image: string }) => void;
}

const UpdateDialog: React.FC<DialogProps> = ({ isOpen, onClose, name, image, onSubmit }) => {
  const [replicas, setReplicas] = useState<number>(0);
  const [deploymentName, setDeploymentName] = useState<string>(name);
  const [deploymentImage, setDeploymentImage] = useState<string>(image);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      name: deploymentName,
      replicas,
      image: deploymentImage,
    });
    onClose(); // Close the dialog after submission
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-75 z-50">
      <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
        <h2 className="text-xl font-semibold mb-4">Update Replica and Image</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="name">
              Name
            </label>
            <input
              type="text"
              id="name"
              value={deploymentName}
              onChange={(e) => setDeploymentName(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
              disabled
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="replicas">
              Replicas
            </label>
            <input
              type="number"
              id="replicas"
              value={replicas}
              onChange={(e) => setReplicas(Number(e.target.value))}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="image">
              Image
            </label>
            <input
              type="text"
              id="image"
              value={deploymentImage}
              onChange={(e) => setDeploymentImage(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
              disabled
            />
          </div>
          <div className="flex justify-end">
            <button
              type="button"
              onClick={onClose}
              className="mr-2 py-2 px-4 bg-gray-500 text-white rounded-md hover:bg-gray-600"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700"
            >
              Update
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default UpdateDialog;