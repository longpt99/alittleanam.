'use client';

import Link from 'next/link';
import Image from 'next/image';
import {
  bellIcon,
  bellOutlineIcon,
  cartIcon,
  cartOutlineIcon,
  heartIcon,
  heartOutlineIcon,
  homeIcon,
  homeOutlineIcon,
  profileIcon,
  profileOutlineIcon,
} from '../../public/icons';
import { usePathname } from 'next/navigation';

const links = [
  {
    name: 'Home',
    href: '/',
    icon: homeOutlineIcon,
    activeIcon: homeIcon,
  },
  {
    name: 'Favorites',
    href: '/favorites',
    icon: heartOutlineIcon,
    activeIcon: heartIcon,
  },
  {
    name: 'Cart',
    href: '/cart',
    icon: cartOutlineIcon,
    activeIcon: cartIcon,
  },
  {
    name: 'Notifications',
    href: '/products/1231',
    icon: bellOutlineIcon,
    activeIcon: bellIcon,
  },
  {
    name: 'Profile',
    href: '/profile',
    icon: profileOutlineIcon,
    activeIcon: profileIcon,
  },
];

export default function NavLinks() {
  const pathname = usePathname();

  const linkItems = links.map((link) => {
    const isActive = pathname === link.href;
    const iconSrc = isActive ? link.activeIcon : link.icon;

    return (
      <li key={link.name} className="ml-6 mr-6">
        <Link href={link.href}>
          <Image width={24} height={24} src={iconSrc} alt={link.name} />
        </Link>
      </li>
    );
  });

  return (
    <div className="fixed bottom-0 w-full mx-auto my-0 bg-white rounded-[20px_20px_0_0] border-[thin] border-solid border-[#F1F1F1]">
      <ul className="flex justify-center items-center h-[74px]">{linkItems}</ul>
    </div>
  );
}
