<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('Parcourir'),
  'category' => t('Charte2016'),
  'icon' => 'parcourir.png',
  'theme' => 'parcourir',
  'admin css' => 'parcourir.css',
  'regions' => array(
    'navigation' => t('Navigation'),
    'block2' => t('Block 2'),
    'block3' => t('Block 3'),
    'block4' => t('Block 4'),
    'block5' => t('Block 5'),
    'bottom' => t('Bottom'),
  ),
);
