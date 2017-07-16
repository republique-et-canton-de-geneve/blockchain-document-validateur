<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('Organisation overview'),
  'category' => t('Charte2016'),
  'icon' => 'organisation-overview.png',
  'theme' => 'organisation-overview',
  'admin css' => 'organisation-overview.css',
  'regions' => array(
    'block1' => t('Block 1'),
    'block2' => t('Block 2'),
    'block3' => t('Block 3'),
    'block4' => t('Block 4'),
    'block5before' => t('Block 5 - before'),
    'block5' => t('Block 5 - grid view'),
    'block6' => t('Block 6'),
    'block7' => t('Block 7'),
  ),
);
