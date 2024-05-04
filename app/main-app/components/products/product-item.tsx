'use client';

import { useEffect, useState } from 'react';
import useEmblaCarousel from 'embla-carousel-react';
import clsx from 'clsx';
import styles from './product-item.module.css';

interface Props {
  title: string;
  properties: string[];
  require: boolean;
}

export default function ProductItem(props: Props) {
  const [emblaRef] = useEmblaCarousel({
    dragFree: true,
  });
  const [activeButtons, setActiveButtons] = useState<boolean[]>(() => {
    const arr = Array(props.properties.length).fill(false);
    if (props.require) {
      arr[0] = true;
    }

    return arr;
  });

  // Function to toggle the active state of a button
  const toggleButton = (index: number) => {
    const updatedButtons = Array(props.properties.length).fill(false);
    if (props.require) {
      updatedButtons[index] = !updatedButtons[index];
    } else {
      updatedButtons[index] = !activeButtons[index];
    }
    setActiveButtons(updatedButtons);
  };

  return (
    <div>
      <h4 className="font-semibold text-[1.2rem] mx-0 my-4">{props.title}</h4>
      {/* <div className="flex justify-between"> */}
      <div className="embla" ref={emblaRef}>
        <div className="embla__container">
          {props.properties.map((item, i) => (
            <button
              key={i}
              className={clsx(
                `px-11 py-2.5 rounded-[10px] border-[thin] border-solid embla__slide text-center mr-2.5 last-child:mt-0`,
                {
                  'text-[color:var(--primary-color)] bg-[#FFF5EE] border-[color:var(--primary-color)]':
                    activeButtons[i],
                  'border-[#DEDEDE]': !activeButtons[i],
                }
              )}
              onClick={() => toggleButton(i)} // Toggle the active state on click
            >
              {item}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}
