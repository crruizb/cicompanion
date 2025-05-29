const SpinnerLoader = () => (
  <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/30">
    <div className="w-10 h-10 border-4 border-gray-200 border-t-neutral-600 rounded-full animate-spin"></div>
  </div>
);

export default SpinnerLoader;
