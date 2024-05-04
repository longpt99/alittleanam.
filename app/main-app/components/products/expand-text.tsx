'use client';

import { useEffect, useRef, useState } from 'react';

export const ExpandableText = ({ children }: any) => {
  const textRef = useRef(null);
  const [isExpanded, setIsExpanded] = useState(false);
  const [isButtonNeeded, setIsButtonNeeded] = useState(false);

  // Toggles the isExpanded state between true and false
  const toggleExpanded = () => {
    setIsExpanded(!isExpanded);
  };

  function checkOverflow(element: any) {
    return element.scrollHeight > maxHeight;
  }

  const lineHeight = 16; // Line height in pixels
  const maxLinesToShow = 3;
  const maxHeight = lineHeight * maxLinesToShow;

  // After your component mounts, check if overflow is happening
  useEffect(() => {
    setIsButtonNeeded(checkOverflow(textRef.current));
  }, []);

  return (
    <>
      <p
        ref={textRef}
        className={`text-[#808080] text-xs ${isExpanded ? '' : 'line-clamp-3'}`}
      >
        {children}
      </p>
      {isButtonNeeded && (
        <button
          className="text-xs text-[color:var(--primary-color)] font-medium hover:underline"
          onClick={toggleExpanded}
        >
          {isExpanded ? 'See Less' : 'See More'}
        </button>
      )}
    </>
  );
};
