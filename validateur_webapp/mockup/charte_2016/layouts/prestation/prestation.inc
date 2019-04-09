<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('prestation'),
  'category' => t('Bootstrap'),
  'category' => t('Charte2016'),
  'icon' => 'prestation.png',
  'theme' => 'prestation',
  'admin css' => 'prestation.css',
  'regions' => array(
    'navigation' => t('Navigation'),
    'navigation_droite' => t('Navigation droite'),
    'block2' => t('Block 2'),
    'block3' => t('Block 3'),
    'block4' => t('Block 4'),
    'block5' => t('Block 5'),
    'bottom' => t('Bottom'),
  ),
);
